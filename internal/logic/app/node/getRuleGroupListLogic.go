package node

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetRuleGroupListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get rule group list
func NewGetRuleGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRuleGroupListLogic {
	return &GetRuleGroupListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRuleGroupListLogic) GetRuleGroupList() (resp *types.AppRuleGroupListResponse, err error) {
	nodeRuleGroupList, err := l.svcCtx.ServerModel.QueryAllRuleGroup(l.ctx)
	if err != nil {
		l.Logger.Error("[GetRuleGroupList] get subscribe rule group list failed: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe rule group list failed: %v", err.Error())
	}
	nodeRuleGroups := make([]types.ServerRuleGroup, 0)
	tool.DeepCopy(&nodeRuleGroups, nodeRuleGroupList)
	return &types.AppRuleGroupListResponse{
		Total: int64(len(nodeRuleGroups)),
		List:  nodeRuleGroups,
	}, nil
}
