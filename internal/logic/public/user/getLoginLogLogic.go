package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
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
