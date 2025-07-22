package payment

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/payment"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// GetPaymentMethodListHandler Get Payment Method List
func GetPaymentMethodListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetPaymentMethodListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := payment.NewGetPaymentMethodListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetPaymentMethodList(&req)
		result.HttpResult(c, resp, err)
	}
}
