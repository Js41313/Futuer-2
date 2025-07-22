package coupon

import (
	"context"
	"math/rand"
	"time"

	"github.com/Js41313/Futuer-2/internal/model/coupon"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/random"
	"github.com/Js41313/Futuer-2/pkg/snowflake"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateCouponLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create coupon
func NewCreateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCouponLogic {
	return &CreateCouponLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCouponLogic) CreateCoupon(req *types.CreateCouponRequest) error {
	if req.Code == "" {
		rand.NewSource(time.Now().UnixNano())
		sid := snowflake.GetID()
		req.Code = random.KeyNew(4, 2) + "-" + random.StrToDashedString(random.EncodeBase36(sid))
	}
	couponInfo := &coupon.Coupon{}
	tool.DeepCopy(couponInfo, req)
	couponInfo.Subscribe = tool.Int64SliceToString(req.Subscribe)
	err := l.svcCtx.CouponModel.Insert(l.ctx, couponInfo)
	if err != nil {
		l.Errorw("[CreateCoupon] Database Error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create coupon error: %v", err.Error())
	}
	return nil
}
