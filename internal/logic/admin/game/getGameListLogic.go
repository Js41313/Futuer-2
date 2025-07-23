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

type GetGameListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGameListLogic {
	return &GetGameListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGameListLogic) GetGameList(req *types.GetGameListRequest) (*types.GetGameListResponse, error) {
	filter := &game.GameFilter{
		Search: req.Search,
		Page:   req.Page,
		Size:   req.Size,
	}

	total, list, err := l.svcCtx.GameModel.FindGameListByFilter(l.ctx, filter)
	if err != nil {
		l.Errorw("[GetGameList] Find Database Error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get game list error: %v", err)
	}

	return &types.GetGameListResponse{
		Total: total,
		List:  list,
	}, nil
}
