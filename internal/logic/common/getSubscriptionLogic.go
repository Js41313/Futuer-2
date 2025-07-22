package common

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetSubscriptionLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Subscription
func NewGetSubscriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscriptionLogic {
	return &GetSubscriptionLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscriptionLogic) GetSubscription() (resp *types.GetSubscriptionResponse, err error) {
	resp = &types.GetSubscriptionResponse{
		List: make([]types.Subscribe, 0),
	}
	// Get the subscription list
	data, err := l.svcCtx.SubscribeModel.QuerySubscribeListByShow(l.ctx)
	if err != nil {
		l.Errorw("[Site GetSubscription]", logger.Field("err", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscription list error: %v", err.Error())
	}
	tool.DeepCopy(&resp.List, data)
	return
}
