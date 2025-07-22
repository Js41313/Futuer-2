package middleware

import (
	"strings"

	"github.com/Js41313/Futuer-2/internal/logic/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/gin-gonic/gin"
)

func PanDomainMiddleware(svc *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		if svc.Config.Subscribe.PanDomain {
			domain := c.Request.Host
			domainArr := strings.Split(domain, ".")
			domainFirst := domainArr[0]
			request := types.SubscribeRequest{
				Token: domainFirst,
				Flag:  domainArr[1],
				UA:    c.Request.Header.Get("User-Agent"),
			}
			l := subscribe.NewSubscribeLogic(c, svc)
			resp, err := l.Generate(&request)
			if err != nil {
				return
			}
			c.Header("subscription-userinfo", resp.Header)
			c.String(200, "%s", string(resp.Config))
			c.Abort()
			return
		}
		c.Next()
	}
}
