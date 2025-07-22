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

type GetNodeConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeConfigLogic {
	return &GetNodeConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeConfigLogic) GetNodeConfig() (*types.NodeConfig, error) {
	resp := &types.NodeConfig{}

	// get server config from db
	configs, err := l.svcCtx.SystemModel.GetNodeConfig(l.ctx)
	if err != nil {
		l.Errorw("[GetNodeConfigLogic] GetNodeConfig get server config error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetNodeConfig get server config error: %v", err.Error())
	}
	// reflect to response
	tool.SystemConfigSliceReflectToStruct(configs, resp)
	return resp, nil
}
