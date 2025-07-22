package server

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetNodeGroupListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeGroupListLogic {
	return &GetNodeGroupListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeGroupListLogic) GetNodeGroupList() (resp *types.GetNodeGroupListResponse, err error) {
	nodeGroupList, err := l.svcCtx.ServerModel.QueryAllGroup(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), err.Error())
	}
	nodeGroups := make([]types.ServerGroup, 0)
	tool.DeepCopy(&nodeGroups, nodeGroupList)
	return &types.GetNodeGroupListResponse{
		Total: int64(len(nodeGroups)),
		List:  nodeGroups,
	}, nil
}
