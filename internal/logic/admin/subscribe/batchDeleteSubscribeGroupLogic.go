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

type BatchDeleteSubscribeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Batch delete subscribe group
func NewBatchDeleteSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteSubscribeGroupLogic {
	return &BatchDeleteSubscribeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteSubscribeGroupLogic) BatchDeleteSubscribeGroup(req *types.BatchDeleteSubscribeGroupRequest) error {
	err := l.svcCtx.DB.Model(&subscribe.Group{}).Where("id in ?", req.Ids).Delete(&subscribe.Group{}).Error
	if err != nil {
		l.Logger.Error("[BatchDeleteSubscribeGroup] Delete Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "batch delete subscribe group failed: %v", err.Error())
	}
	return nil
}
