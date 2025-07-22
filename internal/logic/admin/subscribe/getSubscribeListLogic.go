package subscribe

import (
	"context"
	"encoding/json"

	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
)

type GetSubscribeListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get subscribe list
func NewGetSubscribeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeListLogic {
	return &GetSubscribeListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeListLogic) GetSubscribeList(req *types.GetSubscribeListRequest) (resp *types.GetSubscribeListResponse, err error) {
	total, list, err := l.svcCtx.SubscribeModel.QuerySubscribeListByPage(l.ctx, int(req.Page), int(req.Size), req.GroupId, req.Search)
	if err != nil {
		l.Logger.Error("[GetSubscribeListLogic] get subscribe list failed: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe list failed: %v", err.Error())
	}
	var (
		subscribeIdList = make([]int64, 0, len(list))
		resultList      = make([]types.SubscribeItem, 0, len(list))
	)
	for _, item := range list {
		subscribeIdList = append(subscribeIdList, item.Id)
		var sub types.SubscribeItem
		tool.DeepCopy(&sub, item)
		if item.Discount != "" {
			err = json.Unmarshal([]byte(item.Discount), &sub.Discount)
			if err != nil {
				l.Logger.Error("[GetSubscribeListLogic] JSON unmarshal failed: ", logger.Field("error", err.Error()), logger.Field("discount", item.Discount))
			}
		}
		sub.Server = tool.StringToInt64Slice(item.Server)
		sub.ServerGroup = tool.StringToInt64Slice(item.ServerGroup)
		resultList = append(resultList, sub)
	}

	subscribeMaps, err := l.svcCtx.UserModel.QueryActiveSubscriptions(l.ctx, subscribeIdList...)
	if err != nil {
		l.Logger.Error("[GetSubscribeListLogic] get user subscribe failed: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get user subscribe failed: %v", err.Error())
	}

	for i, item := range resultList {
		if subscribe, ok := subscribeMaps[item.Id]; ok {
			resultList[i].Sold = subscribe
		}
	}

	resp = &types.GetSubscribeListResponse{
		Total: total,
		List:  resultList,
	}
	return
}
