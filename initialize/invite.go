package initialize

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

func Invite(ctx *svc.ServiceContext) {
	// Initialize the system configuration
	logger.Debug("Register config initialization")
	configs, err := ctx.SystemModel.GetInviteConfig(context.Background())
	if err != nil {
		logger.Error("[Init Invite Config] Get Invite Config Error: ", logger.Field("error", err.Error()))
		return
	}
	var inviteConfig config.InviteConfig
	tool.SystemConfigSliceReflectToStruct(configs, &inviteConfig)
	ctx.Config.Invite = inviteConfig
}
