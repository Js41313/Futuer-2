package server

import (
	"github.com/Js41313/Futuer-2/internal/logic/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Push online users
func PushOnlineUsersHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.OnlineUsersRequest
		_ = c.ShouldBind(&req)
		_ = c.ShouldBindQuery(&req.ServerCommon)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := server.NewPushOnlineUsersLogic(c.Request.Context(), svcCtx)
		err := l.PushOnlineUsers(&req)
		result.HttpResult(c, nil, err)
	}
}
