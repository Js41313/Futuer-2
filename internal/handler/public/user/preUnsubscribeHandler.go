package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Pre Unsubscribe
func PreUnsubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.PreUnsubscribeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewPreUnsubscribeLogic(c.Request.Context(), svcCtx)
		resp, err := l.PreUnsubscribe(&req)
		result.HttpResult(c, resp, err)
	}
}
