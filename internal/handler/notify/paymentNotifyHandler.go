package notify

import (
	"fmt"
	"net/http"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/logic/notify"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/payment"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// PaymentNotifyHandler Payment Notify
func PaymentNotifyHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		platform, ok := c.Request.Context().Value(constant.CtxKeyPlatform).(string)
		if !ok {
			logger.WithContext(c.Request.Context()).Errorf("platform not found")
			result.HttpResult(c, nil, fmt.Errorf("platform not found"))
			return
		}

		switch payment.ParsePlatform(platform) {
		case payment.EPay:
			req := &types.EPayNotifyRequest{}
			if err := c.ShouldBind(req); err != nil {
				result.HttpResult(c, nil, err)
				return
			}
			l := notify.NewEPayNotifyLogic(c, svcCtx)
			if err := l.EPayNotify(req); err != nil {
				logger.WithContext(c.Request.Context()).Errorf("EPayNotify failed: %v", err.Error())
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			c.String(http.StatusOK, "%s", "success")
		case payment.Stripe:
			l := notify.NewStripeNotifyLogic(c.Request.Context(), svcCtx)
			if err := l.StripeNotify(c.Request, c.Writer); err != nil {
				result.HttpResult(c, nil, err)
				return
			}
			result.HttpResult(c, nil, nil)

		case payment.AlipayF2F:
			l := notify.NewAlipayNotifyLogic(c.Request.Context(), svcCtx)
			if err := l.AlipayNotify(c.Request); err != nil {
				result.HttpResult(c, nil, err)
				return
			}
			// Return success to alipay
			c.String(http.StatusOK, "%s", "success")

		default:
			logger.WithContext(c.Request.Context()).Errorf("platform %s not support", platform)
		}
	}
}
