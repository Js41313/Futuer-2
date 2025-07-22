package oauth

import (
	"net/http"

	"github.com/Js41313/Futuer-2/internal/logic/auth/oauth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Apple Login Callback
func AppleLoginCallbackHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.AppleLoginCallbackRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}
		l := oauth.NewAppleLoginCallbackLogic(c, svcCtx)
		err := l.AppleLoginCallback(&req, c.Request, c.Writer)
		if err != nil {
			result.HttpResult(c, nil, err)
		}
	}
}
