package system

import (
	"context"

	"github.com/Js41313/Futuer-2/initialize"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetVerifyConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVerifyConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerifyConfigLogic {
	return &GetVerifyConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVerifyConfigLogic) GetVerifyConfig() (*types.VerifyConfig, error) {
	resp := &types.VerifyConfig{}
	// get verify config from db
	verifyConfigs, err := l.svcCtx.SystemModel.GetVerifyConfig(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get verify config failed: %v", err.Error())
	}
	// reflect to response
	tool.SystemConfigSliceReflectToStruct(verifyConfigs, resp)
	// update verify config to system
	initialize.Verify(l.svcCtx)
	return resp, nil
}
