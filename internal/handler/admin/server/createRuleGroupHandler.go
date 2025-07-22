package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// CreateRuleGroupHandler Create rule group
func CreateRuleGroupHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CreateRuleGroupRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := server.NewCreateRuleGroupLogic(c.Request.Context(), svcCtx)
		err := l.CreateRuleGroup(&req)
		result.HttpResult(c, nil, err)
	}
}
