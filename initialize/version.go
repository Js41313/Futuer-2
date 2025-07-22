package initialize

import (
	"errors"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/initialize/migrate"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/orm"
)

func Migrate(ctx *svc.ServiceContext) {
	mc := orm.Mysql{
		Config: ctx.Config.MySQL,
	}
	if err := migrate.Migrate(mc.Dsn()).Up(); err != nil {
		if errors.Is(err, migrate.NoChange) {
			logger.Info("[Migrate] database not change")
			return
		}
		logger.Errorf("[Migrate] Up error: %v", err.Error())
		panic(err)
	}
	// if not found admin user
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&user.User{}).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			if err := migrate.CreateAdminUser(ctx.Config.Administrator.Email, ctx.Config.Administrator.Password, tx); err != nil {
				logger.Errorf("[Migrate] CreateAdminUser error: %v", err.Error())
				return err
			}
			logger.Info("[Migrate] Create admin user success")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
