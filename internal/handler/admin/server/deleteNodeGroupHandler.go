package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Delete node group
func DeleteNodeGroupHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.DeleteNodeGroupRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := server.NewDeleteNodeGroupLogic(c.Request.Context(), svcCtx)
		err := l.DeleteNodeGroup(&req)
		result.HttpResult(c, nil, err)
	}
}
