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

type GetGlobalConfigLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get global config
func NewGetGlobalConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGlobalConfigLogic {
	return &GetGlobalConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGlobalConfigLogic) GetGlobalConfig() (resp *types.GetGlobalConfigResponse, err error) {
	resp = new(types.GetGlobalConfigResponse)

	currencyCfg, err := l.svcCtx.SystemModel.GetCurrencyConfig(l.ctx)
	if err != nil {
		l.Logger.Error("[GetGlobalConfigLogic] GetCurrencyConfig error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetCurrencyConfig error: %v", err.Error())
	}
	verifyCodeCfg, err := l.svcCtx.SystemModel.GetVerifyCodeConfig(l.ctx)
	if err != nil {
		l.Logger.Error("[GetGlobalConfigLogic] GetVerifyCodeConfig error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetVerifyCodeConfig error: %v", err.Error())
	}

	tool.DeepCopy(&resp.Site, l.svcCtx.Config.Site)
	tool.DeepCopy(&resp.Subscribe, l.svcCtx.Config.Subscribe)
	tool.DeepCopy(&resp.Auth.Email, l.svcCtx.Config.Email)
	tool.DeepCopy(&resp.Auth.Mobile, l.svcCtx.Config.Mobile)
	tool.DeepCopy(&resp.Auth.Register, l.svcCtx.Config.Register)
	tool.DeepCopy(&resp.Verify, l.svcCtx.Config.Verify)
	tool.DeepCopy(&resp.Invite, l.svcCtx.Config.Invite)
	tool.SystemConfigSliceReflectToStruct(currencyCfg, &resp.Currency)
	tool.SystemConfigSliceReflectToStruct(verifyCodeCfg, &resp.VerifyCode)

	resp.Verify = types.VeifyConfig{
		TurnstileSiteKey:          l.svcCtx.Config.Verify.TurnstileSiteKey,
		EnableLoginVerify:         l.svcCtx.Config.Verify.LoginVerify,
		EnableRegisterVerify:      l.svcCtx.Config.Verify.RegisterVerify,
		EnableResetPasswordVerify: l.svcCtx.Config.Verify.ResetPasswordVerify,
	}
	var methods []string

	// auth methods
	authMethods, err := l.svcCtx.AuthModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Error("[GetGlobalConfigLogic] FindAll error: ", logger.Field("error", err.Error()))
	}

	for _, method := range authMethods {
		if *method.Enabled {
			methods = append(methods, method.Method)
		}
	}
	resp.OAuthMethods = methods

	webAds, err := l.svcCtx.SystemModel.FindOneByKey(l.ctx, "WebAD")
	if err != nil {
		l.Logger.Error("[GetGlobalConfigLogic] FindOneByKey error: ", logger.Field("error", err.Error()), logger.Field("key", "WebAD"))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOneByKey error: %v", err.Error())
	}
	// web ads config
	resp.WebAd = webAds.Value == "true"
	return
}
