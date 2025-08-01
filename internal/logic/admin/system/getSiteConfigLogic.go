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

type GetSiteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger.Logger
}

func NewGetSiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteConfigLogic {
	return &GetSiteConfigLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSiteConfigLogic) GetSiteConfig() (resp *types.SiteConfig, err error) {
	resp = &types.SiteConfig{}
	// get site config from db
	siteConfigs, err := l.svcCtx.SystemModel.GetSiteConfig(l.ctx)
	if err != nil {
		l.Logger.Error("[GetSiteConfig] Database query error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get site config failed: %v", err.Error())
	}
	// reflect to response
	tool.SystemConfigSliceReflectToStruct(siteConfigs, resp)
	return resp, nil
}
