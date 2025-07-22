package auth

import (
	"github.com/Js41313/Futuer-2/internal/logic/auth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Check user telephone is exist
func CheckUserTelephoneHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.TelephoneCheckUserRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := auth.NewCheckUserTelephoneLogic(c.Request.Context(), svcCtx)
		resp, err := l.CheckUserTelephone(&req)
		result.HttpResult(c, resp, err)
	}
}
