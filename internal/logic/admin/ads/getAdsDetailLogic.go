package ads

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetAdsDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Ads Detail
func NewGetAdsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdsDetailLogic {
	return &GetAdsDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdsDetailLogic) GetAdsDetail(req *types.GetAdsDetailRequest) (resp *types.Ads, err error) {
	data, err := l.svcCtx.AdsModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Errorw("find ads error", logger.Field("error", err.Error()), logger.Field("id", req.Id))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find ads error: %v", err.Error())
	}
	resp = new(types.Ads)
	tool.DeepCopy(resp, data)
	return
}
