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

type GetCurrencyConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Currency Config
func NewGetCurrencyConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrencyConfigLogic {
	return &GetCurrencyConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrencyConfigLogic) GetCurrencyConfig() (resp *types.CurrencyConfig, err error) {
	configs, err := l.svcCtx.SystemModel.GetCurrencyConfig(l.ctx)
	if err != nil {
		l.Errorw("[GetCurrencyConfigLogic] GetCurrencyConfig error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetCurrencyConfig error: %v", err.Error())
	}
	resp = &types.CurrencyConfig{}
	tool.SystemConfigSliceReflectToStruct(configs, resp)
	return
}
