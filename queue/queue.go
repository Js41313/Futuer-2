package queue

import (
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/queue/handler"
	"github.com/hibiken/asynq"
)

type Service struct {
	svc    *svc.ServiceContext
	server *asynq.Server
}

func NewService(svc *svc.ServiceContext) *Service {
	return &Service{
		svc:    svc,
		server: initService(svc),
	}
}

func (m *Service) Start() {
	logger.Infof("start consumer service")
	mux := asynq.NewServeMux()
	// register tasks
	handler.RegisterHandlers(mux, m.svc)
	if err := m.server.Run(mux); err != nil {
		logger.Error("consumer service error", logger.LogField{
			Key:   "error",
			Value: err.Error(),
		})
	}
}

func (m *Service) Stop() {
	logger.Info("stop consumer service")
	m.server.Stop()
}

func initService(svc *svc.ServiceContext) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: svc.Config.Redis.Host, Password: svc.Config.Redis.Pass, DB: 5},
		asynq.Config{
			IsFailure: func(err error) bool {
				logger.Error("consumer service error", logger.Field("error", err.Error()))
				return true
			},
			Concurrency: 20,
		},
	)
}
