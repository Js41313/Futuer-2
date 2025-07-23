package game

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/game"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateGameLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGameLogic {
	return &UpdateGameLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGameLogic) UpdateGame(req *types.UpdateGameRequest) error {
	gameData := &game.Game{
		Id:          req.Id,
		Icon:        req.Icon,
		Cover:       req.Cover,
		Name:        req.Name,
		Region:      req.Region,
		Process:     req.Process,
		Route:       req.Route,
		Description: req.Description,
	}

	err := l.svcCtx.GameModel.Update(l.ctx, gameData)
	if err != nil {
		l.Errorw("[UpdateGame] Update Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update game error: %v", err)
	}

	return nil
}
