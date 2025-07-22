package coupon

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/coupon"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Batch delete coupon
func BatchDeleteCouponHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.BatchDeleteCouponRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := coupon.NewBatchDeleteCouponLogic(c.Request.Context(), svcCtx)
		err := l.BatchDeleteCoupon(&req)
		result.HttpResult(c, nil, err)
	}
}
