package user

import (
	"context"

	"github.com/perfect-panel/server/pkg/constant"

	"github.com/perfect-panel/server/internal/model/user"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
)

type GetLoginLogLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Login Log
func NewGetLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogLogic {
	return &GetLoginLogLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogLogic) GetLoginLog(req *types.GetLoginLogRequest) (resp *types.GetLoginLogResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	data, total, err := l.svcCtx.UserModel.FilterLoginLogList(l.ctx, req.Page, req.Size, &user.LoginLogFilterParams{
		UserId: u.Id,
	})
	if err != nil {
		l.Errorw("find login log failed:", logger.Field("error", err.Error()), logger.Field("user_id", u.Id))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find login log failed: %v", err.Error())
	}
	list := make([]types.UserLoginLog, 0)
	tool.DeepCopy(&list, data)
	return &types.GetLoginLogResponse{
		Total: total,
		List:  list,
	}, nil
}
