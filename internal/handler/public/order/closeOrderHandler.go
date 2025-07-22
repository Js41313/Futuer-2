package order

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/order"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Close order
func CloseOrderHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CloseOrderRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := order.NewCloseOrderLogic(c.Request.Context(), svcCtx)
		err := l.CloseOrder(&req)
		result.HttpResult(c, nil, err)
	}
}
