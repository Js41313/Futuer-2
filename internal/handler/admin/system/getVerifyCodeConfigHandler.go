package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Verify Code Config
func GetVerifyCodeConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := system.NewGetVerifyCodeConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetVerifyCodeConfig()
		result.HttpResult(c, resp, err)
	}
}
