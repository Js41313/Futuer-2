package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update subscribe
func UpdateSubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateSubscribeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := subscribe.NewUpdateSubscribeLogic(c.Request.Context(), svcCtx)
		err := l.UpdateSubscribe(&req)
		result.HttpResult(c, nil, err)
	}
}
