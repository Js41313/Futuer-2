package game

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteGameLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGameLogic {
	return &DeleteGameLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGameLogic) DeleteGame(req *types.DeleteGameRequest) error {
	err := l.svcCtx.GameModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[DeleteGame] Delete Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete game error: %v", err)
	}

	return nil
}
