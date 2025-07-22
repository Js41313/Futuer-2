package authMethod

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/authMethod"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update auth method config
func UpdateAuthMethodConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateAuthMethodConfigRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := authMethod.NewUpdateAuthMethodConfigLogic(c.Request.Context(), svcCtx)
		resp, err := l.UpdateAuthMethodConfig(&req)
		result.HttpResult(c, resp, err)
	}
}
