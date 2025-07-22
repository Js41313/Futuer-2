package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Js41313/Futuer-2/internal/logic/admin/game"
	"github.com/Js41313/Futuer-2/internal/model/game"
	"github.com/Js41313/Futuer-2/internal/svc"
)

var logic *game.GameLogic

func InitGameHandler(svcCtx *svc.ServiceContext) {
	logic = game.NewGameLogic(svcCtx.DB)
}

// GameHandler 游戏管理相关接口
func GameHandler(r *gin.RouterGroup) {
	r.POST("/v1/admin/game", CreateGame)
	r.PUT("/v1/admin/game", UpdateGame)
	r.DELETE("/v1/admin/game", DeleteGame)
	r.GET("/v1/admin/game/list", GetGameList)
	r.GET("/v1/admin/game/detail", GetGameDetail)
}

func CreateGame(c *gin.Context) {
	var req game.Game
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := logic.CreateGame(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CreateGame", "data": req})
}

func UpdateGame(c *gin.Context) {
	var req game.Game
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	if err := logic.UpdateGame(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "UpdateGame", "data": req})
}

func DeleteGame(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	var gameID int64
	if _, err := fmt.Sscan(id, &gameID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := logic.DeleteGame(gameID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DeleteGame", "id": gameID})
}

func GetGameList(c *gin.Context) {
	offset := 0
	limit := 10
	if v := c.Query("offset"); v != "" {
		fmt.Sscan(v, &offset)
	}
	if v := c.Query("limit"); v != "" {
		fmt.Sscan(v, &limit)
	}
	games, total, err := logic.GetGameList(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "GetGameList", "data": games, "total": total})
}

func GetGameDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	var gameID int64
	if _, err := fmt.Sscan(id, &gameID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	gameData, err := logic.GetGameDetail(gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "GetGameDetail", "data": gameData})
}