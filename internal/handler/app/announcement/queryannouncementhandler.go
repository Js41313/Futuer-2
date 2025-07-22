package announcement

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/announcement"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Query announcement
func QueryAnnouncementHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.QueryAnnouncementRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := announcement.NewQueryAnnouncementLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryAnnouncement(&req)
		result.HttpResult(c, resp, err)
	}
}
