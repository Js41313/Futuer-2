package log

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/log"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get message log list
func GetMessageLogListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetMessageLogListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := log.NewGetMessageLogListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetMessageLogList(&req)
		result.HttpResult(c, resp, err)
	}
}
