package order

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/order"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Checkout order
func CheckoutOrderHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CheckoutOrderRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := order.NewCheckoutOrderLogic(c.Request.Context(), svcCtx)
		resp, err := l.CheckoutOrder(&req, c.Request.Host)
		result.HttpResult(c, resp, err)
	}
}
