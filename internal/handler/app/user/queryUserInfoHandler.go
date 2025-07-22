package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// query user info
func QueryUserInfoHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := user.NewQueryUserInfoLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryUserInfo()
		result.HttpResult(c, resp, err)
	}
}
