package order

import (
	"context"
	"encoding/json"
	"time"

	"github.com/perfect-panel/server/pkg/constant"
	"github.com/perfect-panel/server/pkg/xerr"

	"gorm.io/gorm"

	"github.com/hibiken/asynq"
	"github.com/perfect-panel/server/internal/model/order"
	"github.com/perfect-panel/server/internal/model/user"
	"github.com/perfect-panel/server/pkg/tool"
	queue "github.com/perfect-panel/server/queue/types"
	"github.com/pkg/errors"

	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
)

type ResetTrafficLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Reset traffic
func NewResetTrafficLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetTrafficLogic {
	return &ResetTrafficLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetTrafficLogic) ResetTraffic(req *types.ResetTrafficOrderRequest) (resp *types.ResetTrafficOrderResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	// find user subscription
	userSubscribe, err := l.svcCtx.UserModel.FindOneUserSubscribe(l.ctx, req.UserSubscribeID)
	if err != nil {
		l.Error("[ResetTraffic] Database query error", logger.Field("error", err.Error()), logger.Field("UserSubscribeID", req.UserSubscribeID))
		return nil, errors.Wrapf(err, "find user subscribe error: %v", err.Error())
	}
	if userSubscribe.Subscribe == nil {
		l.Error("[ResetTraffic] subscribe not found", logger.Field("UserSubscribeID", req.UserSubscribeID))
		return nil, errors.New("subscribe not found")
	}
	amount := userSubscribe.Subscribe.Replacement
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
	// find payment method
	payment, err := l.svcCtx.PaymentModel.FindOne(l.ctx, req.Payment)
	if err != nil {
		l.Error("[ResetTraffic] Database query error", logger.Field("error", err.Error()), logger.Field("payment", req.Payment))
		return nil, errors.Wrapf(err, "find payment error: %v", err.Error())
	}
	var feeAmount int64
	// Calculate the handling fee
	if amount > 0 {
		feeAmount = calculateFee(amount, payment)
	}
	// create order
	orderInfo := order.Order{
		Id:             0,
		ParentId:       userSubscribe.OrderId,
		UserId:         u.Id,
		OrderNo:        tool.GenerateTradeNo(),
		Type:           3,
		Price:          userSubscribe.Subscribe.Replacement,
		Amount:         amount + feeAmount,
		GiftAmount:     deductionAmount,
		FeeAmount:      feeAmount,
		PaymentId:      req.Payment,
		Method:         payment.Platform,
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
				l.Error("[ResetTraffic] Database update error", logger.Field("error", err.Error()), logger.Field("user", u))
				return err
			}
			// create deduction record
			deductionLog := user.GiftAmountLog{
				UserId:  orderInfo.UserId,
				OrderNo: orderInfo.OrderNo,
				Amount:  orderInfo.GiftAmount,
				Type:    2,
				Balance: u.GiftAmount,
				Remark:  "ResetTraffic order deduction",
			}
			if err := db.Model(&user.GiftAmountLog{}).Create(&deductionLog).Error; err != nil {
				l.Error("[ResetTraffic] Database insert error", logger.Field("error", err.Error()), logger.Field("deductionLog", deductionLog))
				return err
			}
		}
		// insert order
		return db.Model(&order.Order{}).Create(&orderInfo).Error
	})
	if err != nil {
		l.Error("[ResetTraffic] Database insert error", logger.Field("error", err.Error()), logger.Field("order", orderInfo))
		return nil, errors.Wrapf(err, "insert order error: %v", err.Error())
	}
	// Deferred task
	payload := queue.DeferCloseOrderPayload{
		OrderNo: orderInfo.OrderNo,
	}
	val, err := json.Marshal(payload)
	if err != nil {
		l.Error("[ResetTraffic] Marshal payload error", logger.Field("error", err.Error()), logger.Field("payload", payload))
	}
	task := asynq.NewTask(queue.DeferCloseOrder, val, asynq.MaxRetry(3))
	taskInfo, err := l.svcCtx.Queue.Enqueue(task, asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
	if err != nil {
		l.Error("[ResetTraffic] Enqueue task error", logger.Field("error", err.Error()), logger.Field("task", task))
	} else {
		l.Info("[ResetTraffic] Enqueue task success", logger.Field("TaskID", taskInfo.ID))
	}
	return &types.ResetTrafficOrderResponse{
		OrderNo: orderInfo.OrderNo,
	}, nil
}
