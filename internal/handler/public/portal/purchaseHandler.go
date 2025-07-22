package portal

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/portal"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Purchase subscription
func PurchaseHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.PortalPurchaseRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := portal.NewPurchaseLogic(c.Request.Context(), svcCtx)
		resp, err := l.Purchase(&req)
		result.HttpResult(c, resp, err)
	}
}
