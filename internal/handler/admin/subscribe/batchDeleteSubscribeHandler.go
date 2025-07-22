package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Batch delete subscribe
func BatchDeleteSubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.BatchDeleteSubscribeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := subscribe.NewBatchDeleteSubscribeLogic(c.Request.Context(), svcCtx)
		err := l.BatchDeleteSubscribe(&req)
		result.HttpResult(c, nil, err)
	}
}
