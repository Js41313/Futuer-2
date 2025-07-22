package initialize

import (
	"context"
	"encoding/json"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/model/auth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

func Mobile(ctx *svc.ServiceContext) {
	logger.Debug("Mobile config initialization")
	method, err := ctx.AuthModel.FindOneByMethod(context.Background(), "mobile")
	if err != nil {
		panic(err)
	}
	var cfg config.MobileConfig
	var mobileConfig auth.MobileAuthConfig
	mobileConfig.Unmarshal(method.Config)
	tool.DeepCopy(&cfg, mobileConfig)
	cfg.Enable = *method.Enabled
	value, _ := json.Marshal(mobileConfig.PlatformConfig)
	cfg.PlatformConfig = string(value)
	ctx.Config.Mobile = cfg
}
