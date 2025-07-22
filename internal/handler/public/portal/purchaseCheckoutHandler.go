package portal

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/portal"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Purchase Checkout
func PurchaseCheckoutHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CheckoutOrderRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := portal.NewPurchaseCheckoutLogic(c.Request.Context(), svcCtx)
		resp, err := l.PurchaseCheckout(&req)
		result.HttpResult(c, resp, err)
	}
}
