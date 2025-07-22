package game

import (
	"time"

	"gorm.io/gorm"
)

// Game 游戏实体

type Game struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Icon        string    `json:"icon"`
	Cover       string    `json:"cover"`
	Name        string    `json:"name"`
	Region      string    `json:"region"`
	Process     string    `json:"process"`
	Route       string    `json:"route"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateGame 创建游戏
func CreateGame(db *gorm.DB, game *Game) error {
	return db.Create(game).Error
}

// UpdateGame 更新游戏
func UpdateGame(db *gorm.DB, game *Game) error {
	return db.Save(game).Error
}

// DeleteGame 删除游戏
func DeleteGame(db *gorm.DB, id int64) error {
	return db.Delete(&Game{}, id).Error
}

// GetGameList 获取游戏列表
func GetGameList(db *gorm.DB, offset, limit int) ([]Game, int64, error) {
	var games []Game
	var total int64
	db.Model(&Game{}).Count(&total)
	err := db.Offset(offset).Limit(limit).Find(&games).Error
	return games, total, err
}

// GetGameDetail 获取游戏详情
func GetGameDetail(db *gorm.DB, id int64) (*Game, error) {
	var game Game
	err := db.First(&game, id).Error
	return &game, err
}
