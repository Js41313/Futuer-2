package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get user subcribe
func GetUserSubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetUserSubscribeListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewGetUserSubscribeLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetUserSubscribe(&req)
		result.HttpResult(c, resp, err)
	}
}
