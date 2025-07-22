package payment

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/payment"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update Payment Method
func UpdatePaymentMethodHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdatePaymentMethodRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := payment.NewUpdatePaymentMethodLogic(c.Request.Context(), svcCtx)
		resp, err := l.UpdatePaymentMethod(&req)
		result.HttpResult(c, resp, err)
	}
}
