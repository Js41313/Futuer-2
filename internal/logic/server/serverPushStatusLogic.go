package server

import (
	"context"
	"errors"

	"github.com/Js41313/Futuer-2/internal/model/cache"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type ServerPushStatusLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Push server status
func NewServerPushStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServerPushStatusLogic {
	return &ServerPushStatusLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServerPushStatusLogic) ServerPushStatus(req *types.ServerPushStatusRequest) error {
	// Find server info
	serverInfo, err := l.svcCtx.ServerModel.FindOne(l.ctx, req.ServerId)
	if err != nil || serverInfo.Id <= 0 {
		l.Errorw("[PushOnlineUsers] FindOne error", logger.Field("error", err))
		return errors.New("server not found")
	}
	err = l.svcCtx.NodeCache.UpdateNodeStatus(l.ctx, req.ServerId, cache.NodeStatus{
		Cpu:       req.Cpu,
		Mem:       req.Mem,
		Disk:      req.Disk,
		UpdatedAt: req.UpdatedAt,
	})
	if err != nil {
		l.Errorw("[ServerPushStatus] UpdateNodeStatus error", logger.Field("error", err))
		return errors.New("update node status failed")
	}
	return nil
}
