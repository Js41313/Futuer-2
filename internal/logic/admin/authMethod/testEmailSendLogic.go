package authMethod

import (
	"context"
	"fmt"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/email"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type TestEmailSendLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Test email send
func NewTestEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestEmailSendLogic {
	return &TestEmailSendLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestEmailSendLogic) TestEmailSend(req *types.TestEmailSendRequest) error {
	client, err := email.NewSender(l.svcCtx.Config.Email.Platform, l.svcCtx.Config.Email.PlatformConfig, l.svcCtx.Config.Site.SiteName)
	if err != nil {
		l.Errorw("new email sender err", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "new email sender err: %v", err.Error())
	}
	err = client.Send([]string{req.Email}, "Test Email Send", "this a test email send by ppanel")
	if err != nil {
		return errors.Wrapf(xerr.NewErrCodeMsg(500, fmt.Sprintf("send email err: %v", err.Error())), "send email err: %v", err.Error())
	}
	return nil
}
