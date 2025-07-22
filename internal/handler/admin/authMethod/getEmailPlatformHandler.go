package authMethod

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/authMethod"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get email support platform
func GetEmailPlatformHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := authMethod.NewGetEmailPlatformLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetEmailPlatform()
		result.HttpResult(c, resp, err)
	}
}
