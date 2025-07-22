package smslogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/model/log"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/constant"
	"github.com/Js41313/Futuer-2/pkg/sms"
	"github.com/Js41313/Futuer-2/queue/types"
	"github.com/hibiken/asynq"
)

type SmsSendCount struct {
	Count    int   `json:"count"`
	CreateAt int64 `json:"create_at"`
}

type SendSmsLogic struct {
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		svcCtx: svcCtx,
	}
}
func (l *SendSmsLogic) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload types.SendSmsPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Unmarshal payload failed",
			logger.Field("error", err.Error()),
			logger.Field("payload", task.Payload()),
		)
		return nil
	}
	client, err := sms.NewSender(l.svcCtx.Config.Mobile.Platform, l.svcCtx.Config.Mobile.PlatformConfig)
	if err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] New send sms client failed", logger.Field("error", err.Error()), logger.Field("payload", payload))
		return err
	}
	createSms := &log.MessageLog{
		Type:     log.Mobile.String(),
		Platform: l.svcCtx.Config.Mobile.Platform,
		To:       fmt.Sprintf("+%s%s", payload.TelephoneArea, payload.Telephone),
		Subject:  constant.ParseVerifyType(payload.Type).String(),
		Content:  "",
	}
	err = client.SendCode(payload.TelephoneArea, payload.Telephone, payload.Content)

	createSms.Content = client.GetSendCodeContent(payload.Content)

	if err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Send sms failed", logger.Field("error", err.Error()), logger.Field("payload", payload))
		if l.svcCtx.Config.Model != constant.DevMode {
			createSms.Status = 2
		} else {
			return nil
		}
	}
	createSms.Status = 1
	logger.WithContext(ctx).Info("[SendSmsLogic] Send sms", logger.Field("telephone", payload.Telephone), logger.Field("content", createSms.Content))
	err = l.svcCtx.LogModel.InsertMessageLog(ctx, createSms)
	if err != nil {
		logger.WithContext(ctx).Error("[SendSmsLogic] Send sms failed", logger.Field("error", err.Error()), logger.Field("payload", payload))
		return nil
	}
	return nil
}
