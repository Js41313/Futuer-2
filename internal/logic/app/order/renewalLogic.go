package order

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/order"
	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	queue "github.com/Js41313/Futuer-2/queue/types"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RenewalLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Renewal Subscription
func NewRenewalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenewalLogic {
	return &RenewalLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RenewalLogic) Renewal(req *types.RenewalOrderRequest) (resp *types.RenewalOrderResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	orderNo := tool.GenerateTradeNo()
	// find user subscribe
	userSubscribe, err := l.svcCtx.UserModel.FindOneUserSubscribe(l.ctx, req.UserSubscribeID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find user subscribe error: %v", err.Error())
	}
	// find subscription
	sub, err := l.svcCtx.SubscribeModel.FindOne(l.ctx, userSubscribe.SubscribeId)
	if err != nil {
		l.Error("[Renewal] Database query error", logger.Field("error", err.Error()), logger.Field("subscribe_id", userSubscribe.SubscribeId))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find subscribe error: %v", err.Error())
	}
	// check subscribe plan status
	if !*sub.Sell {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "subscribe not sell")
	}
	var discount float64 = 1
	if sub.Discount != "" {
		var dis []types.SubscribeDiscount
		_ = json.Unmarshal([]byte(sub.Discount), &dis)
		discount = getDiscount(dis, req.Quantity)
	}
	price := sub.UnitPrice * req.Quantity
	amount := int64(float64(price) * discount)
	discountAmount := price - amount
	var coupon int64 = 0
	if req.Coupon != "" {
		couponInfo, err := l.svcCtx.CouponModel.FindOneByCode(l.ctx, req.Coupon)
		if err != nil {
			l.Error("[Renewal] Database query error", logger.Field("error", err.Error()), logger.Field("coupon", req.Coupon))
			return nil, errors.Wrapf(err, "find coupon error: %v", err.Error())
		}
		if couponInfo.Count <= couponInfo.UsedCount {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CouponInsufficientUsage), "coupon used")
		}
		coupon = calculateCoupon(amount, couponInfo)
	}
	payment, err := l.svcCtx.PaymentModel.FindOne(l.ctx, req.Payment)
	if err != nil {
		l.Error("[Renewal] Database query error", logger.Field("error", err.Error()), logger.Field("payment", req.Payment))
		return nil, errors.Wrapf(err, "find payment error: %v", err.Error())
	}
	amount -= coupon

	var deductionAmount int64
	// Check user deduction amount
	if u.GiftAmount > 0 {
		if u.GiftAmount >= amount {
			deductionAmount = amount
			amount = 0
			u.GiftAmount -= amount
		} else {
			deductionAmount = u.GiftAmount
			amount -= u.GiftAmount
			u.GiftAmount = 0
		}
	}

	var feeAmount int64
	// Calculate the handling fee
	if amount > 0 {
		feeAmount = calculateFee(amount, payment)
	}

	amount += feeAmount

	// create order
	orderInfo := order.Order{
		UserId:         u.Id,
		ParentId:       userSubscribe.OrderId,
		OrderNo:        orderNo,
		Type:           2,
		Quantity:       req.Quantity,
		Price:          price,
		Amount:         amount,
		GiftAmount:     deductionAmount,
		Discount:       discountAmount,
		Coupon:         req.Coupon,
		CouponDiscount: coupon,
		PaymentId:      payment.Id,
		Method:         payment.Platform,
		FeeAmount:      feeAmount,
		Status:         1,
		SubscribeId:    userSubscribe.SubscribeId,
		SubscribeToken: userSubscribe.Token,
	}
	// Database transaction
	err = l.svcCtx.DB.Transaction(func(db *gorm.DB) error {
		// update user deduction && Pre deduction ,Return after canceling the order
		if orderInfo.GiftAmount > 0 {
			// update user deduction && Pre deduction ,Return after canceling the order
			if err := l.svcCtx.UserModel.Update(l.ctx, u, db); err != nil {
				l.Error("[Purchase] Database update error", logger.Field("error", err.Error()), logger.Field("user", u))
				return err
			}
			// create deduction record
			deductionLog := user.GiftAmountLog{
				UserId:  orderInfo.UserId,
				OrderNo: orderInfo.OrderNo,
				Amount:  orderInfo.GiftAmount,
				Type:    2,
				Balance: u.GiftAmount,
				Remark:  "Renewal order deduction",
			}
			if err := db.Model(&user.GiftAmountLog{}).Create(&deductionLog).Error; err != nil {
				l.Error("[Renewal] Database insert error", logger.Field("error", err.Error()), logger.Field("deductionLog", deductionLog))
				return err
			}
		}
		// insert order
		return db.Model(&order.Order{}).Create(&orderInfo).Error
	})
	if err != nil {
		l.Error("[Renewal] Database insert error", logger.Field("error", err.Error()), logger.Field("order", orderInfo))
		return nil, errors.Wrapf(err, "insert order error: %v", err.Error())
	}
	// Deferred task
	payload := queue.DeferCloseOrderPayload{
		OrderNo: orderInfo.OrderNo,
	}
	val, err := json.Marshal(payload)
	if err != nil {
		l.Error("[Renewal] Marshal payload error", logger.Field("error", err.Error()), logger.Field("payload", payload))
	}
	task := asynq.NewTask(queue.DeferCloseOrder, val, asynq.MaxRetry(3))
	taskInfo, err := l.svcCtx.Queue.Enqueue(task, asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
	if err != nil {
		l.Error("[Renewal] Enqueue task error", logger.Field("error", err.Error()), logger.Field("task", task))
	} else {
		l.Info("[Renewal] Enqueue task success", logger.Field("TaskID", taskInfo.ID))
	}
	return &types.RenewalOrderResponse{
		OrderNo: orderInfo.OrderNo,
	}, nil
}
