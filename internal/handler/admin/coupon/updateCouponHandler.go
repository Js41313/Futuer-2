package coupon

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/coupon"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update coupon
func UpdateCouponHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateCouponRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := coupon.NewUpdateCouponLogic(c.Request.Context(), svcCtx)
		err := l.UpdateCoupon(&req)
		result.HttpResult(c, nil, err)
	}
}
