package payment

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/payment"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get available payment methods
func GetAvailablePaymentMethodsHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := payment.NewGetAvailablePaymentMethodsLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetAvailablePaymentMethods()
		result.HttpResult(c, resp, err)
	}
}
