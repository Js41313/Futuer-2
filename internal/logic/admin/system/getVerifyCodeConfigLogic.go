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

type GetVerifyCodeConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Verify Code Config
func NewGetVerifyCodeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerifyCodeConfigLogic {
	return &GetVerifyCodeConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVerifyCodeConfigLogic) GetVerifyCodeConfig() (resp *types.VerifyCodeConfig, err error) {
	data, err := l.svcCtx.SystemModel.GetVerifyCodeConfig(l.ctx)
	if err != nil {
		l.Errorw("Get Verify Code Config Error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Get Verify Code Config Error: %s", err.Error())
	}
	resp = &types.VerifyCodeConfig{}
	tool.SystemConfigSliceReflectToStruct(data, resp)
	return
}
