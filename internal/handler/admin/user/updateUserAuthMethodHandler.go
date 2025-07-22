package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update user auth method
func UpdateUserAuthMethodHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateUserAuthMethodRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewUpdateUserAuthMethodLogic(c.Request.Context(), svcCtx)
		err := l.UpdateUserAuthMethod(&req)
		result.HttpResult(c, nil, err)
	}
}
