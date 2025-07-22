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

type UpdateSubscribeGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update subscribe group
func NewUpdateSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubscribeGroupLogic {
	return &UpdateSubscribeGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubscribeGroupLogic) UpdateSubscribeGroup(req *types.UpdateSubscribeGroupRequest) error {
	err := l.svcCtx.DB.Model(&subscribe.Group{}).Where("id = ?", req.Id).Save(&subscribe.Group{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}).Error
	if err != nil {
		l.Logger.Error("[UpdateSubscribeGroup] update subscribe group failed", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update subscribe group failed: %v", err.Error())
	}
	return nil
}
