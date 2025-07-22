package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get user subcribe by id
func GetUserSubscribeByIdHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetUserSubscribeByIdRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewGetUserSubscribeByIdLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetUserSubscribeById(&req)
		result.HttpResult(c, resp, err)
	}
}
