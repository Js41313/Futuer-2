package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get subscribe list
func QuerySubscribeListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := subscribe.NewQuerySubscribeListLogic(c.Request.Context(), svcCtx)
		resp, err := l.QuerySubscribeList()
		result.HttpResult(c, resp, err)
	}
}
