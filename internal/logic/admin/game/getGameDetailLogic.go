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

type GetGameDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGameDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGameDetailLogic {
	return &GetGameDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGameDetailLogic) GetGameDetail(req *types.GetGameDetailRequest) (*game.Game, error) {
	gameDetail, err := l.svcCtx.GameModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[GetGameDetail] Find Database Error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get game detail error: %v", err)
	}

	return gameDetail, nil
}
