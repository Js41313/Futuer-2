package handler

import (
	"github.com/Js41313/Futuer-2/internal/logic/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/gin-gonic/gin"
)

func SubscribeHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.SubscribeRequest
		if c.Request.Header.Get("token") != "" {
			req.Token = c.Request.Header.Get("token")
		} else {
			req.Token = c.Query("token")
		}
		req.UA = c.Request.Header.Get("User-Agent")
		req.Flag = c.Query("flag")
		l := subscribe.NewSubscribeLogic(c, svcCtx)
		resp, err := l.Generate(&req)
		if err != nil {
			return
		}
		c.Header("subscription-userinfo", resp.Header)
		c.String(200, "%s", string(resp.Config))
	}
}

func RegisterSubscribeHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	path := serverCtx.Config.Subscribe.SubscribePath
	if path == "" {
		path = "/api/subscribe"
	}
	router.GET(path, SubscribeHandler(serverCtx))
}
