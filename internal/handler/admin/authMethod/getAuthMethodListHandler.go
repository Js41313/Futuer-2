package authMethod

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/authMethod"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get auth method list
func GetAuthMethodListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := authMethod.NewGetAuthMethodListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetAuthMethodList()
		result.HttpResult(c, resp, err)
	}
}
