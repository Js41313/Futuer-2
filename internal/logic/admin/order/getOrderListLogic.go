package order

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetOrderListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetOrderListLogic Get order list
func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListRequest) (resp *types.GetOrderListResponse, err error) {
	total, list, err := l.svcCtx.OrderModel.QueryOrderListByPage(l.ctx, int(req.Page), int(req.Size), req.Status, req.UserId, req.SubscribeId, req.Search)
	if err != nil {
		l.Errorw("[GetOrderList] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryOrderListByPage error: %v", err.Error())
	}
	resp = &types.GetOrderListResponse{}
	resp.List = make([]types.Order, 0)
	tool.DeepCopy(&resp.List, list)
	resp.Total = total
	return
}
