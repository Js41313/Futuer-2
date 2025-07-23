package game

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/game"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

func GetGameDetailHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.GetGameDetailRequest
		if err := c.ShouldBindUri(&req); err != nil {
			result.HttpResult(c, nil, err)
			return
		}

		// 参数验证
		if validateErr := svcCtx.Validate(&req); validateErr != nil {
			result.HttpResult(c, nil, validateErr)
			return
		}

		l := game.NewGetGameDetailLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetGameDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
