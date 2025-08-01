package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetUserSubscribeLogsLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user subcribe logs
func NewGetUserSubscribeLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSubscribeLogsLogic {
	return &GetUserSubscribeLogsLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSubscribeLogsLogic) GetUserSubscribeLogs(req *types.GetUserSubscribeLogsRequest) (resp *types.GetUserSubscribeLogsResponse, err error) {
	data, total, err := l.svcCtx.UserModel.FilterSubscribeLogList(l.ctx, req.Page, req.Size, &user.SubscribeLogFilterParams{
		UserSubscribeId: req.SubscribeId,
		UserId:          req.UserId,
	})

	if err != nil {
		l.Errorw("[GetUserSubscribeLogs] Get User Subscribe Logs Error:", logger.Field("err", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Get User Subscribe Logs Error")
	}
	var list []types.UserSubscribeLog
	tool.DeepCopy(&list, data)

	return &types.GetUserSubscribeLogsResponse{
		List:  list,
		Total: total,
	}, err
}
