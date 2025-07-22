package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get application config
func QueryApplicationConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := subscribe.NewQueryApplicationConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryApplicationConfig()
		result.HttpResult(c, resp, err)
	}
}
