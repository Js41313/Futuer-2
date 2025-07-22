package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Bind Telegram
func BindTelegramHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := user.NewBindTelegramLogic(c.Request.Context(), svcCtx)
		resp, err := l.BindTelegram()
		result.HttpResult(c, resp, err)
	}
}
