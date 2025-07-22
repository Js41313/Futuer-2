package subscribe

import (
	"context"
	"encoding/json"

	"github.com/Js41313/Futuer-2/pkg/device"

	"github.com/Js41313/Futuer-2/internal/model/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateSubscribeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update subscribe
func NewUpdateSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubscribeLogic {
	return &UpdateSubscribeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubscribeLogic) UpdateSubscribe(req *types.UpdateSubscribeRequest) error {
	// Query the database to get the subscribe information
	_, err := l.svcCtx.SubscribeModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Error("[UpdateSubscribe] Database query error", logger.Field("error", err.Error()), logger.Field("subscribe_id", req.Id))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe error: %v", err.Error())
	}
	discount := ""
	if len(req.Discount) > 0 {
		val, _ := json.Marshal(req.Discount)
		discount = string(val)
	}
	sub := &subscribe.Subscribe{
		Id:             req.Id,
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
		Sort:           req.Sort,
		DeductionRatio: req.DeductionRatio,
		AllowDeduction: req.AllowDeduction,
		ResetCycle:     req.ResetCycle,
		RenewalReset:   req.RenewalReset,
	}
	err = l.svcCtx.SubscribeModel.Update(l.ctx, sub)
	if err != nil {
		l.Logger.Error("[UpdateSubscribe] update subscribe failed", logger.Field("error", err.Error()), logger.Field("subscribe", sub))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update subscribe error: %v", err.Error())
	}
	l.svcCtx.DeviceManager.Broadcast(device.SubscribeUpdate)
	return nil
}
