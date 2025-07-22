package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Query User Affiliate List
func QueryUserAffiliateListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.QueryUserAffiliateListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewQueryUserAffiliateListLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryUserAffiliateList(&req)
		result.HttpResult(c, resp, err)
	}
}
