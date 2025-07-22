package handler

import (
	"github.com/Js41313/Futuer-2/internal/handler/notify"
	"github.com/Js41313/Futuer-2/internal/middleware"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/gin-gonic/gin"
)

func RegisterNotifyHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	group := router.Group("/v1/notify/")
	group.Use(middleware.NotifyMiddleware(serverCtx))
	{
		group.Any("/:platform/:token", notify.PaymentNotifyHandler(serverCtx))
	}

}
