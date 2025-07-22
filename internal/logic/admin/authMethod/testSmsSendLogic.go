package authMethod

import (
	"context"
	"fmt"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/sms"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type TestSmsSendLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Test sms send
func NewTestSmsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestSmsSendLogic {
	return &TestSmsSendLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestSmsSendLogic) TestSmsSend(req *types.TestSmsSendRequest) error {
	client, err := sms.NewSender(l.svcCtx.Config.Mobile.Platform, l.svcCtx.Config.Mobile.PlatformConfig)
	if err != nil {
		l.Errorw("new sms sender err", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "new sms sender err: %v", err.Error())
	}
	err = client.SendCode(req.AreaCode, req.Telephone, "123456")
	if err != nil {
		l.Errorw("send sms err", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCodeMsg(500, fmt.Sprintf("send sms err: %v", err.Error())), "send sms err: %v", err.Error())
	}
	return nil
}
