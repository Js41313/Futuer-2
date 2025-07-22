package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update Currency Config
func UpdateCurrencyConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CurrencyConfig
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := system.NewUpdateCurrencyConfigLogic(c.Request.Context(), svcCtx)
		err := l.UpdateCurrencyConfig(&req)
		result.HttpResult(c, nil, err)
	}
}
