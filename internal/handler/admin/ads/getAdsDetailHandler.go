package ads

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/ads"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Ads Detail
func GetAdsDetailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetAdsDetailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		ctx := c.Request.Context()
		l := ads.NewGetAdsDetailLogic(ctx, svcCtx)
		resp, err := l.GetAdsDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
