package server

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

type GetNodeTagListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get node tag list
func NewGetNodeTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeTagListLogic {
	return &GetNodeTagListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeTagListLogic) GetNodeTagList() (resp *types.GetNodeTagListResponse, err error) {
	tags, err := l.svcCtx.ServerModel.FindServerTags(l.ctx)
	return &types.GetNodeTagListResponse{
		Tags: tool.RemoveDuplicateElements(tags...),
	}, nil
}
