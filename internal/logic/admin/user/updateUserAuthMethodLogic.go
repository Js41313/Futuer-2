package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateUserAuthMethodLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update user auth method
func NewUpdateUserAuthMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAuthMethodLogic {
	return &UpdateUserAuthMethodLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAuthMethodLogic) UpdateUserAuthMethod(req *types.UpdateUserAuthMethodRequest) error {
	method, err := l.svcCtx.UserModel.FindUserAuthMethodByPlatform(l.ctx, req.UserId, req.AuthType)
	if err != nil {
		l.Errorw("Get user auth method error", logger.Field("error", err.Error()), logger.Field("userId", req.UserId), logger.Field("authType", req.AuthType))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Get user auth method error: %v", err.Error())
	}
	method.AuthType = req.AuthType
	method.AuthIdentifier = req.AuthIdentifier
	if err = l.svcCtx.UserModel.UpdateUserAuthMethods(l.ctx, method); err != nil {
		l.Errorw("Update user auth method error", logger.Field("error", err.Error()), logger.Field("userId", req.UserId), logger.Field("authType", req.AuthType))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "Update user auth method error: %v", err.Error())
	}
	return nil
}
