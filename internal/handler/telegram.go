package handler

import (
	"github.com/Js41313/Futuer-2/internal/logic/telegram"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RegisterTelegramHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	router.POST("/v1/telegram/webhook", TelegramHandler(serverCtx))
}

func TelegramHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// auth secret
		secret := c.Query("secret")
		if secret != tool.Md5Encode(svcCtx.Config.Telegram.BotToken, false) {
			logger.WithContext(c.Request.Context()).Error("[TelegramHandler] Secret is wrong", logger.Field("request secret", secret), logger.Field("config secret", tool.Md5Encode(svcCtx.Config.Telegram.BotToken, false)), logger.Field("token", svcCtx.Config.Telegram.BotToken))
			c.Abort()
			result.HttpResult(c, nil, nil)
			return
		}
		var request tgbotapi.Update
		if err := c.BindJSON(&request); err != nil {
			logger.WithContext(c.Request.Context()).Error("[TelegramHandler] Failed to bind request", logger.Field("error", err.Error()))
			c.Abort()
			result.HttpResult(c, nil, err)
		}
		l := telegram.NewTelegramLogic(c, svcCtx)
		l.TelegramLogic(&request)
	}
}
