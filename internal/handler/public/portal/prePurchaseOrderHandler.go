package portal

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/portal"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Pre Purchase Order
func PrePurchaseOrderHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.PrePurchaseOrderRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := portal.NewPrePurchaseOrderLogic(c.Request.Context(), svcCtx)
		resp, err := l.PrePurchaseOrder(&req)
		result.HttpResult(c, resp, err)
	}
}
