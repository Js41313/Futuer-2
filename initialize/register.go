package initialize

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

func Register(ctx *svc.ServiceContext) {
	logger.Debug("Register config initialization")
	configs, err := ctx.SystemModel.GetRegisterConfig(context.Background())
	if err != nil {
		logger.Errorf("[Init Register Config] Get Register Config Error: %s", err.Error())
		return
	}
	var registerConfig config.RegisterConfig
	tool.SystemConfigSliceReflectToStruct(configs, &registerConfig)
	ctx.Config.Register = registerConfig
}
