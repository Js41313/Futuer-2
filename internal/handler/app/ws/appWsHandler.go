package ws

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/ws"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// App heartbeat
func AppWsHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Logic: App heartbeat
		l := ws.NewAppWsLogic(ctx, svcCtx)
		err := l.AppWs(c.Writer, c.Request, c.Param("userid"), c.Param("identifier"))
		result.HttpResult(c, nil, err)
	}
}
