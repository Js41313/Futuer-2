package ads

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/ads"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Delete Ads
func DeleteAdsHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.DeleteAdsRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		ctx := c.Request.Context()
		l := ads.NewDeleteAdsLogic(ctx, svcCtx)
		err := l.DeleteAds(&req)
		result.HttpResult(c, nil, err)
	}
}
