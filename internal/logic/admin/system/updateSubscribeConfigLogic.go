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
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UpdateSubscribeConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubscribeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubscribeConfigLogic {
	return &UpdateSubscribeConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubscribeConfigLogic) UpdateSubscribeConfig(req *types.SubscribeConfig) error {
	v := reflect.ValueOf(*req)
	// Get the reflection type of the structure
	t := v.Type()
	err := l.svcCtx.SystemModel.Transaction(l.ctx, func(db *gorm.DB) error {
		var err error
		for i := 0; i < v.NumField(); i++ {
			// Get the field name
			fieldName := t.Field(i).Name
			// Get the field value to string
			fieldValue := tool.ConvertValueToString(v.Field(i))
			// Update the site config
			err = db.Model(&system.System{}).Where("`category` = 'subscribe' and `key` = ?", fieldName).Update("value", fieldValue).Error
			if err != nil {
				break
			}
		}
		return l.svcCtx.Redis.Del(l.ctx, config.SubscribeConfigKey, config.GlobalConfigKey).Err()
	})

	if err != nil {
		l.Errorw("[UpdateSubscribeConfigLogic] update subscribe config error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update subscribe config error: %v", err)
	}

	if l.svcCtx.Config.Subscribe.SubscribePath != req.SubscribePath {
		go func(svc *svc.ServiceContext) {
			err = svc.Restart()
			if err != nil {
				l.Errorw("[UpdateSubscribeConfigLogic] restart error: ", logger.Field("error", err.Error()))
			}
		}(l.svcCtx)
		return nil
	}

	initialize.Subscribe(l.svcCtx)
	return nil
}
