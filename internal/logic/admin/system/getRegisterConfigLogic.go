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

type GetRegisterConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegisterConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegisterConfigLogic {
	return &GetRegisterConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegisterConfigLogic) GetRegisterConfig() (*types.RegisterConfig, error) {
	resp := &types.RegisterConfig{}

	// get register config from database
	configs, err := l.svcCtx.SystemModel.GetRegisterConfig(l.ctx)
	if err != nil {
		l.Errorw("[GetRegisterConfig] Database query error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get register config error: %v", err.Error())
	}

	// reflect to response
	tool.SystemConfigSliceReflectToStruct(configs, resp)
	return resp, nil
}
