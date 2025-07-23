package game

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/game"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

func CreateGameHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateGameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			result.HttpResult(c, nil, err)
			return
		}

		// 参数验证
		if err := svcCtx.Validate(&req); err != nil {
			result.HttpResult(c, nil, err)
			return
		}

		l := game.NewCreateGameLogic(c.Request.Context(), svcCtx)
		err := l.CreateGame(&req)
		result.HttpResult(c, nil, err)
	}
}
