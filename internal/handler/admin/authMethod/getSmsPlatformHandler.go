package authMethod

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/authMethod"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get sms support platform
func GetSmsPlatformHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := authMethod.NewGetSmsPlatformLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSmsPlatform()
		result.HttpResult(c, resp, err)
	}
}
