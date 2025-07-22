package console

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/console"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Query server total data
func QueryServerTotalDataHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := console.NewQueryServerTotalDataLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryServerTotalData()
		result.HttpResult(c, resp, err)
	}
}
