package svc

import (
	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/hibiken/asynq"
)

func NewAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass, DB: 5})
}
