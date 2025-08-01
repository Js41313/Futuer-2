package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteUserSubscribeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDeleteUserSubscribeLogic Delete user subcribe
func NewDeleteUserSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserSubscribeLogic {
	return &DeleteUserSubscribeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserSubscribeLogic) DeleteUserSubscribe(req *types.DeleteUserSubscribeRequest) error {
	err := l.svcCtx.UserModel.DeleteSubscribeById(l.ctx, req.UserSubscribeId)
	if err != nil {
		l.Errorw("failed to delete user subscribe", logger.Field("error", err.Error()), logger.Field("userSubscribeId", req.UserSubscribeId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "failed to delete user subscribe: %v", err.Error())
	}
	return nil
}
