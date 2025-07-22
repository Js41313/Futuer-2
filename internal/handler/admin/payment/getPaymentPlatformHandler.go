package payment

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/payment"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get supported payment platform
func GetPaymentPlatformHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := payment.NewGetPaymentPlatformLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetPaymentPlatform()
		result.HttpResult(c, resp, err)
	}
}
