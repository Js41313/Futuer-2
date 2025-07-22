package server

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BatchDeleteNodeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeleteNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteNodeLogic {
	return &BatchDeleteNodeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteNodeLogic) BatchDeleteNode(req *types.BatchDeleteNodeRequest) error {
	err := l.svcCtx.DB.Transaction(func(db *gorm.DB) error {
		for _, id := range req.Ids {
			err := l.svcCtx.ServerModel.Delete(l.ctx, id)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.Errorw("[BatchDeleteNode] Delete Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), err.Error())
	}
	return nil
}
