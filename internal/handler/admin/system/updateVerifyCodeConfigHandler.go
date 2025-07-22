package system

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update Verify Code Config
func UpdateVerifyCodeConfigHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.VerifyCodeConfig
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := system.NewUpdateVerifyCodeConfigLogic(c.Request.Context(), svcCtx)
		err := l.UpdateVerifyCodeConfig(&req)
		result.HttpResult(c, nil, err)
	}
}
