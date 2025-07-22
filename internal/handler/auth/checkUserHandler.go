package auth

import (
	"github.com/Js41313/Futuer-2/internal/logic/auth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Check user is exist
func CheckUserHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CheckUserRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := auth.NewCheckUserLogic(c.Request.Context(), svcCtx)
		resp, err := l.CheckUser(&req)
		result.HttpResult(c, resp, err)
	}
}
