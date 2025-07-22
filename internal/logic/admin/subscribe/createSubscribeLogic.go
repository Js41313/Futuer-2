package subscribe

import (
	"context"
	"encoding/json"

	"github.com/perfect-panel/server/internal/model/subscribe"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateSubscribeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateSubscribeLogic Create subscribe
func NewCreateSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubscribeLogic {
	return &CreateSubscribeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubscribeLogic) CreateSubscribe(req *types.CreateSubscribeRequest) error {
	discount := ""
	if len(req.Discount) > 0 {
		val, _ := json.Marshal(req.Discount)
		discount = string(val)
	}
	sub := &subscribe.Subscribe{
		Id:             0,
		Name:           req.Name,
		Description:    req.Description,
		UnitPrice:      req.UnitPrice,
		UnitTime:       req.UnitTime,
		Discount:       discount,
		Replacement:    req.Replacement,
		Inventory:      req.Inventory,
		Traffic:        req.Traffic,
		SpeedLimit:     req.SpeedLimit,
		DeviceLimit:    req.DeviceLimit,
		Quota:          req.Quota,
		GroupId:        req.GroupId,
		ServerGroup:    tool.Int64SliceToString(req.ServerGroup),
		Server:         tool.Int64SliceToString(req.Server),
		Show:           req.Show,
		Sell:           req.Sell,
		Sort:           0,
		DeductionRatio: req.DeductionRatio,
		AllowDeduction: req.AllowDeduction,
		ResetCycle:     req.ResetCycle,
		RenewalReset:   req.RenewalReset,
	}
	err := l.svcCtx.SubscribeModel.Insert(l.ctx, sub)
	if err != nil {
		l.Logger.Error("[CreateSubscribeLogic] create subscribe error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create subscribe error: %v", err.Error())
	}

	return nil
}
