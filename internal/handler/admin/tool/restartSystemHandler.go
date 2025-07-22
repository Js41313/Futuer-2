package tool

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/tool"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Restart System
func RestartSystemHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := tool.NewRestartSystemLogic(c.Request.Context(), svcCtx)
		err := l.RestartSystem()
		result.HttpResult(c, nil, err)
	}
}
