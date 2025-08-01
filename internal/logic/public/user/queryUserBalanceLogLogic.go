package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type QueryUserBalanceLogLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Query User Balance Log
func NewQueryUserBalanceLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserBalanceLogLogic {
	return &QueryUserBalanceLogLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserBalanceLogLogic) QueryUserBalanceLog() (resp *types.QueryUserBalanceLogListResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	var data []*user.BalanceLog
	var total int64
	// Query User Balance Log
	err = l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		return db.Model(&user.BalanceLog{}).Order("created_at DESC").Where("user_id = ?", u.Id).Count(&total).Find(&data).Error
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Query User Balance Log failed: %v", err)
	}
	resp = &types.QueryUserBalanceLogListResponse{
		List:  make([]types.UserBalanceLog, 0),
		Total: total,
	}
	tool.DeepCopy(&resp.List, data)
	return
}
