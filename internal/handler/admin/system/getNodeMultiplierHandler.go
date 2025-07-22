package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Node Multiplier
func GetNodeMultiplierHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewGetNodeMultiplierLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetNodeMultiplier()
		result.HttpResult(c, resp, err)
	}
}
