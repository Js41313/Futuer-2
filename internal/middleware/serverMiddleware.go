package middleware

import (
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/gin-gonic/gin"
)

func ServerMiddleware(svc *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		if key, ok := c.GetQuery("secret_key"); ok {
			if key == svc.Config.Node.NodeSecret {
				c.Next()
				return
			}
		}
		c.String(403, "Forbidden")
		c.Abort()
	}
}
