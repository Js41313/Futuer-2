package node

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/node"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get rule group list
func GetRuleGroupListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := node.NewGetRuleGroupListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetRuleGroupList()
		result.HttpResult(c, resp, err)
	}
}
