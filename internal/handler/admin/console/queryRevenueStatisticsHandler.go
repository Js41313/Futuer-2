package console

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/console"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Query revenue statistics
func QueryRevenueStatisticsHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := console.NewQueryRevenueStatisticsLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryRevenueStatistics()
		result.HttpResult(c, resp, err)
	}
}
