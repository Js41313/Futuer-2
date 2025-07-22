package console

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/console"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Query ticket wait reply
func QueryTicketWaitReplyHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := console.NewQueryTicketWaitReplyLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryTicketWaitReply()
		result.HttpResult(c, resp, err)
	}
}
