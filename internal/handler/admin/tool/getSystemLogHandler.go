package tool

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/tool"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get System Log
func GetSystemLogHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := tool.NewGetSystemLogLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSystemLog()
		result.HttpResult(c, resp, err)
	}
}
