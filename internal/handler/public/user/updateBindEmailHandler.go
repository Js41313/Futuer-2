package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update Bind Email
func UpdateBindEmailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateBindEmailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewUpdateBindEmailLogic(c.Request.Context(), svcCtx)
		err := l.UpdateBindEmail(&req)
		result.HttpResult(c, nil, err)
	}
}
