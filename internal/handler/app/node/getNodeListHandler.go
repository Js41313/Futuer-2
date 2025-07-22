package node

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/node"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Node list
func GetNodeListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.AppUserSubscbribeNodeRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := node.NewGetNodeListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetNodeList(&req)
		result.HttpResult(c, resp, err)
	}
}
