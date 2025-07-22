package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get OAuth Methods
func GetOAuthMethodsHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := user.NewGetOAuthMethodsLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetOAuthMethods()
		result.HttpResult(c, resp, err)
	}
}
