package ticket

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/ticket"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Create ticket follow
func CreateTicketFollowHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CreateTicketFollowRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := ticket.NewCreateTicketFollowLogic(c.Request.Context(), svcCtx)
		err := l.CreateTicketFollow(&req)
		result.HttpResult(c, nil, err)
	}
}
