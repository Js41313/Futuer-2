package svc

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/device"

	"github.com/Js41313/Futuer-2/internal/model/ads"
	"github.com/Js41313/Futuer-2/internal/model/cache"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/model/announcement"
	"github.com/Js41313/Futuer-2/internal/model/application"
	"github.com/Js41313/Futuer-2/internal/model/auth"
	"github.com/Js41313/Futuer-2/internal/model/coupon"
	"github.com/Js41313/Futuer-2/internal/model/document"
	"github.com/Js41313/Futuer-2/internal/model/log"
	"github.com/Js41313/Futuer-2/internal/model/order"
	"github.com/Js41313/Futuer-2/internal/model/payment"
	"github.com/Js41313/Futuer-2/internal/model/server"
	"github.com/Js41313/Futuer-2/internal/model/subscribe"
	"github.com/Js41313/Futuer-2/internal/model/subscribeType"
	"github.com/Js41313/Futuer-2/internal/model/system"
	"github.com/Js41313/Futuer-2/internal/model/ticket"
	"github.com/Js41313/Futuer-2/internal/model/traffic"
	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/pkg/limit"
	"github.com/Js41313/Futuer-2/pkg/nodeMultiplier"
	"github.com/Js41313/Futuer-2/pkg/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB                    *gorm.DB
	Redis                 *redis.Client
	Config                config.Config
	Queue                 *asynq.Client
	NodeCache             *cache.NodeCacheClient
	AuthModel             auth.Model
	AdsModel              ads.Model
	LogModel              log.Model
	UserModel             user.Model
	OrderModel            order.Model
	TicketModel           ticket.Model
	ServerModel           server.Model
	SystemModel           system.Model
	CouponModel           coupon.Model
	PaymentModel          payment.Model
	DocumentModel         document.Model
	SubscribeModel        subscribe.Model
	TrafficLogModel       traffic.Model
	ApplicationModel      application.Model
	AnnouncementModel     announcement.Model
	SubscribeTypeModel    subscribeType.Model
	Restart               func() error
	TelegramBot           *tgbotapi.BotAPI
	NodeMultiplierManager *nodeMultiplier.Manager
	AuthLimiter           *limit.PeriodLimit
	DeviceManager         *device.DeviceManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gorm initialize
	db, err := orm.ConnectMysql(orm.Mysql{
		Config: c.MySQL,
	})
	if err != nil {
		panic(err.Error())
	}
	rds := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})
	err = rds.Ping(context.Background()).Err()
	if err != nil {
		panic(err.Error())
	} else {
		_ = rds.FlushAll(context.Background()).Err()
	}
	authLimiter := limit.NewPeriodLimit(86400, 15, rds, config.SendCountLimitKeyPrefix, limit.Align())
	srv := &ServiceContext{
		DB:                db,
		Redis:             rds,
		Config:            c,
		Queue:             NewAsynqClient(c),
		NodeCache:         cache.NewNodeCacheClient(rds),
		AuthLimiter:       authLimiter,
		AdsModel:          ads.NewModel(db, rds),
		LogModel:          log.NewModel(db),
		AuthModel:         auth.NewModel(db, rds),
		UserModel:         user.NewModel(db, rds),
		OrderModel:        order.NewModel(db, rds),
		TicketModel:       ticket.NewModel(db, rds),
		ServerModel:       server.NewModel(db, rds),
		SystemModel:       system.NewModel(db, rds),
		CouponModel:       coupon.NewModel(db, rds),
		PaymentModel:      payment.NewModel(db, rds),
		DocumentModel:     document.NewModel(db, rds),
		SubscribeModel:    subscribe.NewModel(db, rds),
		TrafficLogModel:   traffic.NewModel(db),
		ApplicationModel:  application.NewModel(db, rds),
		AnnouncementModel: announcement.NewModel(db, rds),
	}
	srv.DeviceManager = NewDeviceManager(srv)
	return srv

}
