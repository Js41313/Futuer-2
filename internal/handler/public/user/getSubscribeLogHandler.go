package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Subscribe Log
func GetSubscribeLogHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetSubscribeLogRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewGetSubscribeLogLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSubscribeLog(&req)
		result.HttpResult(c, resp, err)
	}
}
