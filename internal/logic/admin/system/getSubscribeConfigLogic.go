package system

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetSubscribeConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubscribeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeConfigLogic {
	return &GetSubscribeConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeConfigLogic) GetSubscribeConfig() (resp *types.SubscribeConfig, err error) {
	resp = &types.SubscribeConfig{}
	// get subscribe config from db
	subscribeConfigs, err := l.svcCtx.SystemModel.GetSubscribeConfig(l.ctx)
	if err != nil {
		l.Errorw("[GetSubscribeConfig] Database query error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe config failed: %v", err.Error())
	}

	// reflect to response
	tool.SystemConfigSliceReflectToStruct(subscribeConfigs, resp)
	return resp, nil
}
