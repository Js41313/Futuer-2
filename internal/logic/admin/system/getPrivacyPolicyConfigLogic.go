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

type GetPrivacyPolicyConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetPrivacyPolicyConfigLogic get Privacy Policy Config
func NewGetPrivacyPolicyConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrivacyPolicyConfigLogic {
	return &GetPrivacyPolicyConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrivacyPolicyConfigLogic) GetPrivacyPolicyConfig() (resp *types.PrivacyPolicyConfig, err error) {
	resp = &types.PrivacyPolicyConfig{}
	// get tos config from db
	configs, err := l.svcCtx.SystemModel.GetTosConfig(l.ctx)
	if err != nil {
		l.Errorw("[GetTosConfig] GetTosConfig error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetTosConfig error: %v", err.Error())
	}
	// reflect to response
	tool.SystemConfigSliceReflectToStruct(configs, resp)
	return
}
