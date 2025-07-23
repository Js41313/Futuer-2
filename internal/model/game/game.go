package game

import (
	"context"
	"fmt"
	"time"

	"github.com/Js41313/Futuer-2/pkg/cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GameFilter struct {
	Search string
	Page   int
	Size   int
}

type Game struct {
	Id          int64     `gorm:"primary_key" json:"id"`
	Icon        string    `gorm:"type:varchar(255);comment:Game Icon" json:"icon"`
	Cover       string    `gorm:"type:varchar(255);comment:Game Cover" json:"cover"`
	Name        string    `gorm:"type:varchar(100);not null;comment:Game Name" json:"name"`
	Region      string    `gorm:"type:varchar(100);comment:Game Region" json:"region"`
	Process     string    `gorm:"type:varchar(255);comment:Game Process" json:"process"`
	Route       string    `gorm:"type:varchar(255);comment:Game Route" json:"route"`
	Description string    `gorm:"type:text;comment:Game Description" json:"description"`
	CreatedAt   time.Time `gorm:"<-:create;comment:Creation Time" json:"created_at"`
	UpdatedAt   time.Time `gorm:"comment:Update Time" json:"updated_at"`
}

func (Game) TableName() string {
	return "game"
}

type (
	customGameModel struct {
		*defaultGameModel
	}

	defaultGameModel struct {
		cache.CachedConn
		table string
	}

	GameModel interface {
		Insert(ctx context.Context, data *Game) error
		FindOne(ctx context.Context, id int64) (*Game, error)
		Update(ctx context.Context, data *Game) error
		Delete(ctx context.Context, id int64) error
		FindGameListByFilter(ctx context.Context, filter *GameFilter) (total int64, list []*Game, err error)
		BatchDelete(ctx context.Context, ids []int64) error
	}

	customGameLogicModel interface {
		customGameModel
	}
)

var (
	cacheGameIdPrefix = "cache:game:id:"
	cacheGameAllKeys  = "cache:game:all"
)

func newGameModel(conn *gorm.DB, c *redis.Client) *defaultGameModel {
	return &defaultGameModel{
		CachedConn: cache.NewConn(conn, c),
		table:      "game",
	}
}

func NewGameModel(conn *gorm.DB, c *redis.Client) GameModel {
	return &customGameModel{
		defaultGameModel: newGameModel(conn, c),
	}
}

// Insert inserts a game.
func (m *customGameModel) Insert(ctx context.Context, data *Game) error {
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Create(data).Error
	}, cacheGameAllKeys)
}

// FindOne finds a game by id.
func (m *customGameModel) FindOne(ctx context.Context, id int64) (*Game, error) {
	var game Game
	err := m.QueryCtx(ctx, &game, fmt.Sprintf("%s%v", cacheGameIdPrefix, id), func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Game{}).Where("id = ?", id).First(&game).Error
	})
	return &game, err
}

// Update updates a game.
func (m *customGameModel) Update(ctx context.Context, data *Game) error {
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Model(&Game{}).Where("id = ?", data.Id).Updates(data).Error
	}, cacheGameAllKeys, fmt.Sprintf("%s%v", cacheGameIdPrefix, data.Id))
}

// Delete deletes a game.
func (m *customGameModel) Delete(ctx context.Context, id int64) error {
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Where("id = ?", id).Delete(&Game{}).Error
	}, cacheGameAllKeys, fmt.Sprintf("%s%v", cacheGameIdPrefix, id))
}

// FindGameListByFilter finds games by filter.
func (m *customGameModel) FindGameListByFilter(ctx context.Context, filter *GameFilter) (total int64, list []*Game, err error) {
	err = m.QueryNoCacheCtx(ctx, &list, func(conn *gorm.DB, v interface{}) error {
		query := conn.Model(&Game{})
		if filter.Search != "" {
			query = query.Where("name LIKE ?", "%"+filter.Search+"%")
		}
		if err := query.Count(&total).Error; err != nil {
			return err
		}
		offset := (filter.Page - 1) * filter.Size
		return query.Offset(offset).Limit(filter.Size).Order("created_at DESC").Find(v).Error
	})
	return
}

// BatchDelete deletes multiple games.
func (m *customGameModel) BatchDelete(ctx context.Context, ids []int64) error {
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Where("id IN ?", ids).Delete(&Game{}).Error
	}, cacheGameAllKeys)
}
