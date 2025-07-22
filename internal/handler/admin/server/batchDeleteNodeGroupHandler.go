package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Batch delete node group
func BatchDeleteNodeGroupHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.BatchDeleteNodeGroupRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := server.NewBatchDeleteNodeGroupLogic(c.Request.Context(), svcCtx)
		err := l.BatchDeleteNodeGroup(&req)
		result.HttpResult(c, nil, err)
	}
}
