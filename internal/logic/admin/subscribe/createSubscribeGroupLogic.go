package subscribe

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateSubscribeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create subscribe group
func NewCreateSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubscribeGroupLogic {
	return &CreateSubscribeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubscribeGroupLogic) CreateSubscribeGroup(req *types.CreateSubscribeGroupRequest) error {
	err := l.svcCtx.DB.Model(&subscribe.Group{}).Create(&subscribe.Group{
		Name:        req.Name,
		Description: req.Description,
	}).Error
	if err != nil {
		l.Logger.Error("[CreateSubscribeGroupLogic] create subscribe group failed: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create subscribe group failed: %v", err.Error())
	}
	return nil
}
