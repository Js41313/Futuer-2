package system

import (
	"context"
	"reflect"

	"github.com/Js41313/Futuer-2/initialize"
	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/model/system"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UpdateSiteConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSiteConfigLogic {
	return &UpdateSiteConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSiteConfigLogic) UpdateSiteConfig(req *types.SiteConfig) error {
	// Get the reflection value of the structure
	v := reflect.ValueOf(*req)
	// Get the reflection type of the structure
	t := v.Type()
	err := l.svcCtx.SystemModel.Transaction(l.ctx, func(db *gorm.DB) error {
		var err error
		for i := 0; i < v.NumField(); i++ {
			// Get the field name
			fieldName := t.Field(i).Name
			// Get the field value
			fieldValue := v.Field(i)
			err = db.Model(&system.System{}).Where("`category` = 'site' and `key` = ?", fieldName).Update("value", fieldValue.String()).Error
			if err != nil {
				break
			}
		}
		if err != nil {
			return err
		}

		return l.svcCtx.Redis.Del(l.ctx, config.SiteConfigKey, config.GlobalConfigKey).Err()
	})
	if err != nil {
		l.Logger.Error("[UpdateSiteConfig] update site config error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update site config error: %v", err.Error())
	}
	initialize.Site(l.svcCtx)
	return nil
}
