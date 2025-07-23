package game

import (
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/gin-gonic/gin"
)

// GameHandler 游戏管理处理器
func GameHandler(r *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	gameGroup := r.Group("/game")
	{
		gameGroup.POST("/create", CreateGameHandler(svcCtx))
		gameGroup.PUT("/update", UpdateGameHandler(svcCtx))
		gameGroup.DELETE("/delete/:id", DeleteGameHandler(svcCtx))
		gameGroup.GET("/list", GetGameListHandler(svcCtx))
		gameGroup.GET("/detail/:id", GetGameDetailHandler(svcCtx))
		gameGroup.DELETE("/batch-delete", BatchDeleteGameHandler(svcCtx))
	}
}
