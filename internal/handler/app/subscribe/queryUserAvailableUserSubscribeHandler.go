package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Available subscriptions for users
func QueryUserAvailableUserSubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.AppUserSubscribeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := subscribe.NewQueryUserAvailableUserSubscribeLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryUserAvailableUserSubscribe(&req)
		result.HttpResult(c, resp, err)
	}
}
