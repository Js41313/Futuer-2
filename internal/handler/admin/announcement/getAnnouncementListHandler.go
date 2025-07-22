package announcement

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/announcement"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get announcement list
func GetAnnouncementListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetAnnouncementListRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := announcement.NewGetAnnouncementListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetAnnouncementList(&req)
		result.HttpResult(c, resp, err)
	}
}
