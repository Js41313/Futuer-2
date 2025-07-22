package order

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/order"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get order list
func GetOrderListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetOrderListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := order.NewGetOrderListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetOrderList(&req)
		result.HttpResult(c, resp, err)
	}
}
