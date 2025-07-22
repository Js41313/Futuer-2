package ads

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/ads"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetAdsListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Ads List
func NewGetAdsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdsListLogic {
	return &GetAdsListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdsListLogic) GetAdsList(req *types.GetAdsListRequest) (resp *types.GetAdsListResponse, err error) {
	total, data, err := l.svcCtx.AdsModel.GetAdsListByPage(l.ctx, req.Page, req.Size, ads.Filter{
		Search: req.Search,
		Status: req.Status,
	})
	if err != nil {
		l.Errorw("get ads list error", logger.Field("error", err.Error()), logger.Field("req", req))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get ads list error: %v", err.Error())
	}
	resp = &types.GetAdsListResponse{
		Total: total,
		List:  make([]types.Ads, len(data)),
	}
	tool.DeepCopy(&resp.List, data)
	return
}
