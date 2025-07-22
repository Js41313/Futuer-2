package subscribe

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get subscribe group list
func GetSubscribeGroupListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := subscribe.NewGetSubscribeGroupListLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSubscribeGroupList()
		result.HttpResult(c, resp, err)
	}
}
