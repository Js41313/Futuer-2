package initialize

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

func Site(ctx *svc.ServiceContext) {
	logger.Debug("initialize site config")
	configs, err := ctx.SystemModel.GetSiteConfig(context.Background())
	if err != nil {
		panic(err)
	}
	var siteConfig config.SiteConfig
	tool.SystemConfigSliceReflectToStruct(configs, &siteConfig)
	ctx.Config.Site = siteConfig
}
