package auth

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/auth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// GetAppConfig
func GetAppConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.AppConfigRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := auth.NewGetAppConfigLogic(c, svcCtx)
		resp, err := l.GetAppConfig(&req)
		result.HttpResult(c, resp, err)
	}
}
