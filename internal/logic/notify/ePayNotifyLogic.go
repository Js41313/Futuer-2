package notify

import (
	"encoding/json"
	"net/url"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"

	"github.com/Js41313/Futuer-2/internal/model/payment"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/payment/epay"
	"github.com/hibiken/asynq"

	queueType "github.com/Js41313/Futuer-2/queue/types"
)

type EPayNotifyLogic struct {
	logger.Logger
	ctx    *gin.Context
	svcCtx *svc.ServiceContext
}

// EPay notify
func NewEPayNotifyLogic(ctx *gin.Context, svcCtx *svc.ServiceContext) *EPayNotifyLogic {
	return &EPayNotifyLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EPayNotifyLogic) EPayNotify(req *types.EPayNotifyRequest) error {

	// Find payment config
	data, ok := l.ctx.Request.Context().Value(constant.CtxKeyPayment).(*payment.Payment)
	if !ok {
		l.Logger.Error("[EPayNotify] Payment not found in context")
		return errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "payment config not found")
	}
	l.Infof("[EPayNotify] Payment config: %+v", data)
	orderInfo, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, req.OutTradeNo)
	if err != nil {
		l.Logger.Error("[EPayNotify] Find order failed", logger.Field("error", err.Error()), logger.Field("orderNo", req.OutTradeNo))
		return errors.Wrapf(xerr.NewErrCode(xerr.OrderNotExist), "order not exist: %v", req.OutTradeNo)
	}

	var config payment.EPayConfig
	if err := json.Unmarshal([]byte(data.Config), &config); err != nil {
		l.Logger.Errorw("[EPayNotify] Unmarshal config failed", logger.Field("error", err.Error()))
		return err
	}
	// Verify sign
	client := epay.NewClient(config.Pid, config.Url, config.Key)
	if !client.VerifySign(urlParamsToMap(l.ctx.Request.URL.RawQuery)) && !l.svcCtx.Config.Debug {
		l.Logger.Error("[EPayNotify] Verify sign failed")
		return nil
	}
	if req.TradeStatus != "TRADE_SUCCESS" {
		l.Logger.Error("[EPayNotify] Trade status is not success", logger.Field("orderNo", req.OutTradeNo), logger.Field("tradeStatus", req.TradeStatus))
		return nil
	}
	if orderInfo.Status == 5 {
		return nil
	}
	// Update order status
	err = l.svcCtx.OrderModel.UpdateOrderStatus(l.ctx, req.OutTradeNo, 2)
	if err != nil {
		l.Logger.Error("[EPayNotify] Update order status failed", logger.Field("error", err.Error()), logger.Field("orderNo", req.OutTradeNo))
		return err
	}
	// Create activate order task
	payload := queueType.ForthwithActivateOrderPayload{
		OrderNo: req.OutTradeNo,
	}
	bytes, err := json.Marshal(&payload)
	if err != nil {
		l.Logger.Error("[EPayNotify] Marshal payload failed", logger.Field("error", err.Error()))
		return err
	}
	task := asynq.NewTask(queueType.ForthwithActivateOrder, bytes)
	taskInfo, err := l.svcCtx.Queue.EnqueueContext(l.ctx, task)
	if err != nil {
		l.Logger.Error("[EPayNotify] Enqueue task failed", logger.Field("error", err.Error()))
		return err
	}
	l.Logger.Info("[EPayNotify] Enqueue task success", logger.Field("taskInfo", taskInfo))
	return nil
}

func urlParamsToMap(query string) map[string]string {
	params := make(map[string]string)
	values, _ := url.ParseQuery(query)
	for k, v := range values {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	return params
}
