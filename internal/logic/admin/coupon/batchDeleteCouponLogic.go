package coupon

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type BatchDeleteCouponLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Batch delete coupon
func NewBatchDeleteCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteCouponLogic {
	return &BatchDeleteCouponLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteCouponLogic) BatchDeleteCoupon(req *types.BatchDeleteCouponRequest) error {
	// batch delete coupon by ids
	err := l.svcCtx.CouponModel.BatchDelete(l.ctx, req.Ids)
	if err != nil {
		l.Errorw("[BatchDeleteCoupon] Database Error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "batch delete coupon error: %v", err.Error())
	}
	return nil
}
