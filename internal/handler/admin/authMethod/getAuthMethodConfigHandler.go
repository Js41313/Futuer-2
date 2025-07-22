package authMethod

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/authMethod"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get auth method config
func GetAuthMethodConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetAuthMethodConfigRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := authMethod.NewGetAuthMethodConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetAuthMethodConfig(&req)
		result.HttpResult(c, resp, err)
	}
}
