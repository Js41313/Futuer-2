package user

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Update user notify setting
func UpdateUserNotifySettingHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UpdateUserNotifySettingRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := user.NewUpdateUserNotifySettingLogic(c.Request.Context(), svcCtx)
		err := l.UpdateUserNotifySetting(&req)
		result.HttpResult(c, nil, err)
	}
}
