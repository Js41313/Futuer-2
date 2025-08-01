package authMethod

import (
	"context"
	"encoding/json"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetAuthMethodConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetAuthMethodConfigLogic Get auth method config
func NewGetAuthMethodConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthMethodConfigLogic {
	return &GetAuthMethodConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthMethodConfigLogic) GetAuthMethodConfig(req *types.GetAuthMethodConfigRequest) (resp *types.AuthMethodConfig, err error) {
	method, err := l.svcCtx.AuthModel.FindOneByMethod(l.ctx, req.Method)
	if err != nil {
		l.Errorw("find one by method failed", logger.Field("method", req.Method), logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find one by method failed: %v", err.Error())
	}

	resp = new(types.AuthMethodConfig)
	tool.DeepCopy(resp, method)
	if method.Config != "" {
		if err := json.Unmarshal([]byte(method.Config), &resp.Config); err != nil {
			l.Errorw("unmarshal config failed", logger.Field("config", method.Config), logger.Field("error", err.Error()))
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "unmarshal apple config failed: %v", err.Error())
		}
	}
	return
}
