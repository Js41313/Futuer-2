package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Delete subscribe group
func DeleteSubscribeGroupHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.DeleteSubscribeGroupRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := subscribe.NewDeleteSubscribeGroupLogic(c.Request.Context(), svcCtx)
		err := l.DeleteSubscribeGroup(&req)
		result.HttpResult(c, nil, err)
	}
}
