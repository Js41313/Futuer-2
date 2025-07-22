package oauth

import (
	"github.com/Js41313/Futuer-2/internal/logic/auth/oauth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// OAuth login
func OAuthLoginHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.OAthLoginRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := oauth.NewOAuthLoginLogic(c.Request.Context(), svcCtx)
		resp, err := l.OAuthLogin(&req)
		result.HttpResult(c, resp, err)
	}
}
