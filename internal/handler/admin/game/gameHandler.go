package game

import (
	"net/http"

	"github.com/Js41313/Futuer-2/internal/logic/admin/game"
	logicgame "github.com/Js41313/Futuer-2/internal/logic/admin/game"
	modelgame "github.com/Js41313/Futuer-2/internal/model/game"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/gin-gonic/gin"
)

// GameHandler 游戏管理相关接口
func GameHandler(r *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	logic := game.NewGameLogic(svcCtx.DB)
	r.POST("/", CreateGame(logic))
	r.PUT("/", UpdateGame(logic))
	r.DELETE("/", DeleteGame(logic))
	r.GET("/list", GetGameList(logic))
	r.GET("/detail", GetGameDetail(logic))
}

func CreateGame(logic *logicgame.GameLogic) gin.HandlerFunc {
	return func(c *gin.Context) {
	var req struct {
		Icon        string `json:"icon"`
		Cover       string `json:"cover"`
		Name        string `json:"name"`
		Region      string `json:"region"`
		Process     string `json:"process"`
		Route       string `json:"route"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameObj := modelgame.Game{
		Icon:        req.Icon,
		Cover:       req.Cover,
		Name:        req.Name,
		Region:      req.Region,
		Process:     req.Process,
		Route:       req.Route,
		Description: req.Description,
	}
		if err := logic.CreateGame(&gameObj); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "CreateGame", "data": gameObj})
	}
}

func UpdateGame(logic *logicgame.GameLogic) gin.HandlerFunc {
	return func(c *gin.Context) {
	var req struct {
		Id          int64  `json:"id"`
		Icon        string `json:"icon"`
		Cover       string `json:"cover"`
		Name        string `json:"name"`
		Region      string `json:"region"`
		Process     string `json:"process"`
		Route       string `json:"route"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	gameObj := modelgame.Game{
		ID:          req.Id,
		Icon:        req.Icon,
		Cover:       req.Cover,
		Name:        req.Name,
		Region:      req.Region,
		Process:     req.Process,
		Route:       req.Route,
		Description: req.Description,
	}
		if err := logic.UpdateGame(&gameObj); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "UpdateGame", "data": gameObj})
	}
}

func DeleteGame(logic *logicgame.GameLogic) gin.HandlerFunc {
	return func(c *gin.Context) {
	var req struct {
		Id int64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	if err := logic.DeleteGame(req.Id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "DeleteGame"})
	}
}

func GetGameList(logic *logicgame.GameLogic) gin.HandlerFunc {
	return func(c *gin.Context) {
	var req struct {
		Page   int    `form:"page"`
		Size   int    `form:"size"`
		Search string `form:"search"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		req.Page = 1
		req.Size = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	offset := (req.Page - 1) * req.Size
		games, total, err := logic.GetGameList(offset, req.Size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "GetGameList", "data": gin.H{"total": total, "list": games}})
	}
}

func GetGameDetail(logic *logicgame.GameLogic) gin.HandlerFunc {
	return func(c *gin.Context) {
	var req struct {
		Id int64 `form:"id"`
	}
	if err := c.ShouldBindQuery(&req); err != nil || req.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
		gameData, err := logic.GetGameDetail(req.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "GetGameDetail", "data": gameData})
	}
}
