package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get node config
func GetNodeConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewGetNodeConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetNodeConfig()
		result.HttpResult(c, resp, err)
	}
}
