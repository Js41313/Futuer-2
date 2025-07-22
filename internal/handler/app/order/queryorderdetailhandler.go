package order

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/order"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get order
func QueryOrderDetailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.QueryOrderDetailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := order.NewQueryOrderDetailLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryOrderDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
