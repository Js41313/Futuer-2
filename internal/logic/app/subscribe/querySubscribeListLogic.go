package subscribe

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

type QuerySubscribeListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get subscribe list
func NewQuerySubscribeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubscribeListLogic {
	return &QuerySubscribeListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubscribeListLogic) QuerySubscribeList() (resp *types.QuerySubscribeListResponse, err error) {

	data, err := l.svcCtx.SubscribeModel.QuerySubscribeList(l.ctx)
	if err != nil {
		l.Errorw("[QuerySubscribeListLogic] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QuerySubscribeList error: %v", err.Error())
	}
	resp = &types.QuerySubscribeListResponse{
		List:  make([]types.Subscribe, 0),
		Total: int64(len(data)),
	}
	for _, v := range data {
		var sub types.Subscribe
		tool.DeepCopy(&sub, v)
		if v.Discount != "" {
			if err = json.Unmarshal([]byte(v.Discount), &sub.Discount); err != nil {
				l.Errorw("[QuerySubscribeListLogic] json.Unmarshal Error", logger.Field("error", err.Error()), logger.Field("value", v.Discount))
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "json.Unmarshal error: %v", err.Error())
			}
		} else {
			sub.Discount = make([]types.SubscribeDiscount, 0)
		}
		resp.List = append(resp.List, sub)
	}
	return
}
