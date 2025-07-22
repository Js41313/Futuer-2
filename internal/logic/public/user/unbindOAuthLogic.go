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

type UnbindOAuthLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Unbind OAuth
func NewUnbindOAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindOAuthLogic {
	return &UnbindOAuthLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindOAuthLogic) UnbindOAuth(req *types.UnbindOAuthRequest) error {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	if !l.validator(req) {
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidParams), "invalid parameter")
	}
	err := l.svcCtx.UserModel.DeleteUserAuthMethods(l.ctx, u.Id, req.Method)
	if err != nil {
		l.Errorw("delete user auth methods failed:", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete user auth methods failed: %v", err.Error())
	}
	return nil
}
func (l *UnbindOAuthLogic) validator(req *types.UnbindOAuthRequest) bool {
	return req.Method != "" && req.Method != "email" && req.Method != "mobile"
}
