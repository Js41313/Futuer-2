package subscribe

import (
	"github.com/gin-gonic/gin"
	"github.com/perfect-panel/server/internal/logic/app/subscribe"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/pkg/result"
)

// Get subscribe list
func QuerySubscribeListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := subscribe.NewQuerySubscribeListLogic(c.Request.Context(), svcCtx)
		resp, err := l.QuerySubscribeList()
		result.HttpResult(c, resp, err)
	}
}
