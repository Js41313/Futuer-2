package ticket

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/ticket"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get ticket detail
func GetTicketHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetTicketRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := ticket.NewGetTicketLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetTicket(&req)
		result.HttpResult(c, resp, err)
	}
}
