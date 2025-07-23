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

type CreateGameLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGameLogic {
	return &CreateGameLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGameLogic) CreateGame(req *types.CreateGameRequest) error {
	gameData := &game.Game{
		Icon:        req.Icon,
		Cover:       req.Cover,
		Name:        req.Name,
		Region:      req.Region,
		Process:     req.Process,
		Route:       req.Route,
		Description: req.Description,
	}

	err := l.svcCtx.GameModel.Insert(l.ctx, gameData)
	if err != nil {
		l.Errorw("[CreateGame] Insert Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create game error: %v", err)
	}

	return nil
}
