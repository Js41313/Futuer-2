package portal

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/portal"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Subscription
func GetSubscriptionHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		l := portal.NewGetSubscriptionLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetSubscription()
		result.HttpResult(c, resp, err)
	}
}
