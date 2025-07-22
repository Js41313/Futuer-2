package server

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteNodeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNodeGroupLogic {
	return &DeleteNodeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNodeGroupLogic) DeleteNodeGroup(req *types.DeleteNodeGroupRequest) error {
	// Check if the group is empty
	count, err := l.svcCtx.ServerModel.QueryServerCountByServerGroups(l.ctx, []int64{req.Id})
	if err != nil {
		l.Errorw("[DeleteNodeGroup] Query Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "query server error: %v", err)
	}
	if count > 0 {
		return errors.Wrapf(xerr.NewErrCode(xerr.NodeGroupNotEmpty), "group is not empty")
	}
	// Delete the group
	err = l.svcCtx.ServerModel.DeleteGroup(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[DeleteNodeGroup] Delete Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), err.Error())
	}
	return nil
}
