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

type QueryOrderDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get order
func NewQueryOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDetailLogic {
	return &QueryOrderDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderDetailLogic) QueryOrderDetail(req *types.QueryOrderDetailRequest) (resp *types.OrderDetail, err error) {
	orderInfo, err := l.svcCtx.OrderModel.FindOneDetailsByOrderNo(l.ctx, req.OrderNo)
	if err != nil {
		l.Error("[QueryOrderDetail] Database query error", logger.Field("error", err.Error()), logger.Field("order_no", req.OrderNo))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find order error: %v", err.Error())
	}
	resp = &types.OrderDetail{}
	tool.DeepCopy(resp, orderInfo)
	// Prevent commission amount leakage
	resp.Commission = 0
	return
}
