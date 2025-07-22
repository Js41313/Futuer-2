package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdatePasswordLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update Password
func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordRequeset) error {
	userInfo := l.ctx.Value(constant.CtxKeyUser).(*user.User)

	// Verify password
	if !tool.VerifyPassWord(req.Password, userInfo.Password) {
		return errors.Wrapf(xerr.NewErrCode(xerr.UserPasswordError), "user password")
	}
	userInfo.Password = tool.EncodePassWord(req.NewPassword)
	err := l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		l.Errorw("update user password error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update user password")
	}
	return err
}
