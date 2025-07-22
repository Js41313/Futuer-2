package auth

import (
	"time"

	"github.com/Js41313/Futuer-2/internal/logic/auth"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/Js41313/Futuer-2/pkg/turnstile"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Reset password
func TelephoneResetPasswordHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.TelephoneResetPasswordRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}
		// get client ip
		req.IP = c.ClientIP()
		if svcCtx.Config.Verify.ResetPasswordVerify {
			verifyTurns := turnstile.New(turnstile.Config{
				Secret:  svcCtx.Config.Verify.TurnstileSecret,
				Timeout: 3 * time.Second,
			})
			if verify, err := verifyTurns.Verify(c.Request.Context(), req.CfToken, req.IP); err != nil || !verify {
				err = errors.Wrapf(xerr.NewErrCode(xerr.TooManyRequests), "error: %v, verify: %v", err, verify)
				result.HttpResult(c, nil, err)
				return
			}
		}
		l := auth.NewTelephoneResetPasswordLogic(c, svcCtx)
		resp, err := l.TelephoneResetPassword(&req)
		result.HttpResult(c, resp, err)
	}
}
