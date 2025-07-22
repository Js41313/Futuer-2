package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get node detail
func GetNodeDetailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetDetailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := server.NewGetNodeDetailLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetNodeDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
