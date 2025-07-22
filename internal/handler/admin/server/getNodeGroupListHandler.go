package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get node group list
func GetNodeGroupListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := server.NewGetNodeGroupListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetNodeGroupList()
		result.HttpResult(c, resp, err)
	}
}
