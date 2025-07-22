package order

import (
	"context"

	"github.com/perfect-panel/server/pkg/constant"

	"github.com/perfect-panel/server/internal/model/user"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
)

type QueryOrderListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get order list
func NewQueryOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderListLogic {
	return &QueryOrderListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderListLogic) QueryOrderList(req *types.QueryOrderListRequest) (resp *types.QueryOrderListResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	total, data, err := l.svcCtx.OrderModel.QueryOrderListByPage(l.ctx, req.Page, req.Size, 0, u.Id, 0, "")
	if err != nil {
		l.Error("[QueryOrderListLogic] Query order list failed", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Query order list failed")
	}
	resp = &types.QueryOrderListResponse{
		Total: total,
		List:  make([]types.OrderDetail, 0),
	}
	for _, item := range data {
		var orderInfo types.OrderDetail
		tool.DeepCopy(&orderInfo, item)
		// Prevent commission amount leakage
		orderInfo.Commission = 0
		resp.List = append(resp.List, orderInfo)
	}

	return
}
