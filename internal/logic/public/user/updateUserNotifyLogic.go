package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateUserNotifyLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update User Notify
func NewUpdateUserNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserNotifyLogic {
	return &UpdateUserNotifyLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserNotifyLogic) UpdateUserNotify(req *types.UpdateUserNotifyRequest) error {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	if u.Id == 0 {
		return errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "user not login")
	}
	u.EnableLoginNotify = req.EnableLoginNotify
	u.EnableBalanceNotify = req.EnableBalanceNotify
	u.EnableSubscribeNotify = req.EnableSubscribeNotify
	u.EnableTradeNotify = req.EnableTradeNotify
	if err := l.svcCtx.UserModel.Update(l.ctx, u); err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "update user notify error: %v", err.Error())
	}
	return nil
}
