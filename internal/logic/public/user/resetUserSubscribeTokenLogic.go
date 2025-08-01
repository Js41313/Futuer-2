package user

import (
	"context"
	"time"

	"github.com/Js41313/Futuer-2/internal/model/order"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/uuidx"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ResetUserSubscribeTokenLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewResetUserSubscribeTokenLogic Reset User Subscribe Token
func NewResetUserSubscribeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetUserSubscribeTokenLogic {
	return &ResetUserSubscribeTokenLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetUserSubscribeTokenLogic) ResetUserSubscribeToken(req *types.ResetUserSubscribeTokenRequest) error {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	userSub, err := l.svcCtx.UserModel.FindOneUserSubscribe(l.ctx, req.UserSubscribeId)
	if err != nil {
		l.Errorw("FindOneUserSubscribe failed:", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOneUserSubscribe failed: %v", err.Error())
	}
	if userSub.UserId != u.Id {
		l.Errorw("UserSubscribeId does not belong to the current user")
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "UserSubscribeId does not belong to the current user")
	}

	var orderDetails *order.Details
	// find order
	if userSub.OrderId != 0 {
		orderDetails, err = l.svcCtx.OrderModel.FindOneDetails(l.ctx, userSub.OrderId)
		if err != nil {
			l.Errorw("FindOneDetails failed:", logger.Field("error", err.Error()))
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOneDetails failed: %v", err.Error())
		}
	} else {
		// if order id is 0, this a admin create user subscribe
		orderDetails = &order.Details{}
	}

	userSub.Token = uuidx.SubscribeToken(orderDetails.OrderNo + time.Now().Format("20060102150405.000"))
	userSub.UUID = uuid.New().String()
	var newSub user.Subscribe
	tool.DeepCopy(&newSub, userSub)

	err = l.svcCtx.UserModel.UpdateSubscribe(l.ctx, &newSub)
	if err != nil {
		l.Errorw("UpdateSubscribe failed:", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "UpdateSubscribe failed: %v", err.Error())
	}
	return nil
}
