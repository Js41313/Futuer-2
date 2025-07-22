package common

import (
	"github.com/Js41313/Futuer-2/internal/logic/common"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get verification code
func SendEmailCodeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.SendCodeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := common.NewSendEmailCodeLogic(c.Request.Context(), svcCtx)
		resp, err := l.SendEmailCode(&req)
		result.HttpResult(c, resp, err)
	}
}
