package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type DeleteUserDeviceLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete user device
func NewDeleteUserDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserDeviceLogic {
	return &DeleteUserDeviceLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserDeviceLogic) DeleteUserDevice(req *types.DeleteUserDeivceRequest) error {
	err := l.svcCtx.UserModel.DeleteDevice(l.ctx, req.Id)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete user error: %v", err.Error())
	}
	return nil
}
