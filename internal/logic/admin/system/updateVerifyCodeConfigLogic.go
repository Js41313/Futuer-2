package system

import (
	"context"
	"reflect"

	"github.com/Js41313/Futuer-2/internal/config"

	"github.com/Js41313/Futuer-2/internal/model/system"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type UpdateVerifyCodeConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update Verify Code Config
func NewUpdateVerifyCodeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVerifyCodeConfigLogic {
	return &UpdateVerifyCodeConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVerifyCodeConfigLogic) UpdateVerifyCodeConfig(req *types.VerifyCodeConfig) error {
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
			err = db.Model(&system.System{}).Where("`category` = 'verify_code' and `key` = ?", fieldName).Update("value", fieldValue).Error
			if err != nil {
				break
			}
		}
		err = l.svcCtx.Redis.Del(l.ctx, config.VerifyCodeConfigKey).Err()
		return err
	})
	if err != nil {
		l.Errorw("[UpdateRegisterConfig] update verify code config error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update register config error: %v", err.Error())
	}
	return nil
}
