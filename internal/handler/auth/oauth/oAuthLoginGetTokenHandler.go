package oauth

import (
	"github.com/Js41313/Futuer-2/internal/logic/auth/oauth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// OAuth login get token
func OAuthLoginGetTokenHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.OAuthLoginGetTokenRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := oauth.NewOAuthLoginGetTokenLogic(c.Request.Context(), svcCtx)
		resp, err := l.OAuthLoginGetToken(&req, c.ClientIP(), c.Request.UserAgent())
		result.HttpResult(c, resp, err)
	}
}
