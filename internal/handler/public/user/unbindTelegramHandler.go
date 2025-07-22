package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Unbind Telegram
func UnbindTelegramHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := user.NewUnbindTelegramLogic(c.Request.Context(), svcCtx)
		err := l.UnbindTelegram()
		result.HttpResult(c, nil, err)
	}
}
