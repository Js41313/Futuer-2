package server

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateNodeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNodeGroupLogic {
	return &UpdateNodeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNodeGroupLogic) UpdateNodeGroup(req *types.UpdateNodeGroupRequest) error {
	// check server group exist
	nodeGroup, err := l.svcCtx.ServerModel.FindOneGroup(l.ctx, req.Id)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), err.Error())
	}
	nodeGroup.Name = req.Name
	nodeGroup.Description = req.Description
	err = l.svcCtx.ServerModel.UpdateGroup(l.ctx, nodeGroup)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), err.Error())
	}
	return nil
}
