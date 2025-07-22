package subscribe

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteSubscribeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete subscribe
func NewDeleteSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubscribeLogic {
	return &DeleteSubscribeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSubscribeLogic) DeleteSubscribe(req *types.DeleteSubscribeRequest) error {
	// Check if the subscribe exists
	var total int64
	err := l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		return db.Model(&user.Subscribe{}).Where("subscribe_id = ? AND `status` = ?", req.Id, 1).Count(&total).Find(&user.Subscribe{}).Error
	})
	if err != nil {
		l.Logger.Error("[DeleteSubscribeLogic] check subscribe failed: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "check subscribe failed: %v", err.Error())
	}
	if total != 0 {
		return errors.Wrapf(xerr.NewErrCode(xerr.SubscribeIsUsedError), "subscribe is used")
	}

	err = l.svcCtx.SubscribeModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Logger.Error("[DeleteSubscribeLogic] delete subscribe failed: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete subscribe failed: %v", err.Error())
	}
	return nil
}
