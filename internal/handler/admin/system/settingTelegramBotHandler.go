package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// setting telegram bot
func SettingTelegramBotHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewSettingTelegramBotLogic(c.Request.Context(), svcCtx)
		err := l.SettingTelegramBot()
		result.HttpResult(c, nil, err)
	}
}
