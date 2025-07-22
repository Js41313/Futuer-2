package game

import (
	"github.com/Js41313/Futuer-2/internal/model/game"
	"gorm.io/gorm"
)

// GameLogic 游戏管理业务逻辑

type GameLogic struct {
	DB *gorm.DB
}

func NewGameLogic(db *gorm.DB) *GameLogic {
	return &GameLogic{DB: db}
}

func (l *GameLogic) CreateGame(g *game.Game) error {
	return game.CreateGame(l.DB, g)
}

func (l *GameLogic) UpdateGame(g *game.Game) error {
	return game.UpdateGame(l.DB, g)
}

func (l *GameLogic) DeleteGame(id int64) error {
	return game.DeleteGame(l.DB, id)
}

func (l *GameLogic) GetGameList(offset, limit int) ([]game.Game, int64, error) {
	return game.GetGameList(l.DB, offset, limit)
}

func (l *GameLogic) GetGameDetail(id int64) (*game.Game, error) {
	return game.GetGameDetail(l.DB, id)
}