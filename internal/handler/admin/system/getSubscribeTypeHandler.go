package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get subscribe type
func GetSubscribeTypeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewGetSubscribeTypeLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSubscribeType()
		result.HttpResult(c, resp, err)
	}
}
