package initialize

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

func Subscribe(svc *svc.ServiceContext) {
	logger.Debug("Subscribe config initialization")
	configs, err := svc.SystemModel.GetSubscribeConfig(context.Background())
	if err != nil {
		logger.Error("[Init Subscribe Config] Get Subscribe Config Error: ", logger.Field("error", err.Error()))
		return
	}

	var subscribeConfig config.SubscribeConfig
	tool.SystemConfigSliceReflectToStruct(configs, &subscribeConfig)
	svc.Config.Subscribe = subscribeConfig
}
