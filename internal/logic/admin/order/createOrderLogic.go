package order

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/order"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateOrderLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create order
func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) error {
	paymentMethod, err := l.svcCtx.PaymentModel.FindOne(l.ctx, req.PaymentId)
	if err != nil {
		l.Logger.Error("[CreateOrder] PaymentMethod Not Found", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.PaymentMethodNotFound), "PaymentMethod not found: %v", err.Error())
	}

	err = l.svcCtx.OrderModel.Insert(l.ctx, &order.Order{
		UserId:         req.UserId,
		OrderNo:        tool.GenerateTradeNo(),
		Type:           req.Type,
		Quantity:       req.Quantity,
		Price:          req.Price,
		Amount:         req.Amount,
		Discount:       req.Discount,
		Coupon:         req.Coupon,
		CouponDiscount: req.CouponDiscount,
		PaymentId:      req.PaymentId,
		Method:         paymentMethod.Token,
		FeeAmount:      req.FeeAmount,
		TradeNo:        req.TradeNo,
		Status:         req.Status,
		SubscribeId:    req.SubscribeId,
	})
	if err != nil {
		l.Logger.Error("[CreateOrder] Database Error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "Insert error: %v", err.Error())
	}
	return nil
}
