package ticket

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/ticket"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Create ticket
func CreateUserTicketHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CreateUserTicketRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := ticket.NewCreateUserTicketLogic(c.Request.Context(), svcCtx)
		err := l.CreateUserTicket(&req)
		result.HttpResult(c, nil, err)
	}
}
