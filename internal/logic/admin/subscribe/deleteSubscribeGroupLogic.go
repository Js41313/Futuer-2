package subscribe

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteSubscribeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete subscribe group
func NewDeleteSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubscribeGroupLogic {
	return &DeleteSubscribeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSubscribeGroupLogic) DeleteSubscribeGroup(req *types.DeleteSubscribeGroupRequest) error {
	err := l.svcCtx.DB.Model(&subscribe.Group{}).Where("id = ?", req.Id).Delete(&subscribe.Group{}).Error
	if err != nil {
		l.Logger.Error("[DeleteSubscribeGroupLogic] delete subscribe group failed: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete subscribe group failed: %v", err.Error())
	}
	return nil
}
