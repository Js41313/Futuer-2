package ticket

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/ticket"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get ticket list
func GetTicketListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetTicketListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := ticket.NewGetTicketListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetTicketList(&req)
		result.HttpResult(c, resp, err)
	}
}
