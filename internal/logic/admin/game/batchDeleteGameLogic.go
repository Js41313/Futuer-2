package game

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type BatchDeleteGameLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeleteGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteGameLogic {
	return &BatchDeleteGameLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteGameLogic) BatchDeleteGame(req *types.BatchDeleteGameRequest) error {
	err := l.svcCtx.GameModel.BatchDelete(l.ctx, req.Ids)
	if err != nil {
		l.Errorw("[BatchDeleteGame] Delete Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "batch delete game error: %v", err)
	}

	return nil
}
