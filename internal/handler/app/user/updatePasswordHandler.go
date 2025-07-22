package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update Password
func UpdatePasswordHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdatePasswordRequeset
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewUpdatePasswordLogic(c.Request.Context(), svcCtx)
		err := l.UpdatePassword(&req)
		result.HttpResult(c, nil, err)
	}
}
