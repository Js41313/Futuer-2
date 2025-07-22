package coupon

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/coupon"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get coupon list
func GetCouponListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetCouponListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := coupon.NewGetCouponListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetCouponList(&req)
		result.HttpResult(c, resp, err)
	}
}
