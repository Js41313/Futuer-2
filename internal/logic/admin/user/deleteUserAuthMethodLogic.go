package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteUserAuthMethodLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete user auth method
func NewDeleteUserAuthMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAuthMethodLogic {
	return &DeleteUserAuthMethodLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserAuthMethodLogic) DeleteUserAuthMethod(req *types.DeleteUserAuthMethodRequest) error {
	err := l.svcCtx.UserModel.DeleteUserAuthMethods(l.ctx, req.UserId, req.AuthType)
	if err != nil {
		l.Errorw("[DeleteUserAuthMethodLogic] Delete User Auth Method Error:", logger.Field("err", err.Error()), logger.Field("userId", req.UserId), logger.Field("authType", req.AuthType))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "Delete User Auth Method Error")
	}
	return nil
}
