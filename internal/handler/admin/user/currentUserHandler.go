package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Current user
func CurrentUserHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		l := user.NewCurrentUserLogic(c.Request.Context(), svcCtx)
		resp, err := l.CurrentUser()
		result.HttpResult(c, resp, err)
	}
}
