package game

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/game"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

func BatchDeleteGameHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.BatchDeleteGameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			result.HttpResult(c, nil, err)
			return
		}

		// 参数验证
		if validateErr := svcCtx.Validate(&req); validateErr != nil {
			result.HttpResult(c, nil, validateErr)
			return
		}

		l := game.NewBatchDeleteGameLogic(c.Request.Context(), svcCtx)
		err := l.BatchDeleteGame(&req)
		result.HttpResult(c, nil, err)
	}
}
