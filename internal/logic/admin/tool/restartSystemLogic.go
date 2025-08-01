package tool

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type RestartSystemLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Restart System
func NewRestartSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartSystemLogic {
	return &RestartSystemLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestartSystemLogic) RestartSystem() error {
	l.Logger.Info("[RestartSystem]", logger.Field("info", "Restarting system"))
	go func() {
		err := l.svcCtx.Restart()
		if err != nil {
			l.Errorw("[RestartSystem]", logger.Field("error", err.Error()))
		}
		l.Logger.Info("[RestartSystem]", logger.Field("info", "System restarted"))
	}()
	return nil
}
