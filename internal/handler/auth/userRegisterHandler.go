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

// User register
func UserRegisterHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.UserRegisterRequest
		_ = c.ShouldBind(&req)
		// get client ip
		req.IP = c.ClientIP()
		if svcCtx.Config.Verify.RegisterVerify {
			verifyTurns := turnstile.New(turnstile.Config{
				Secret:  svcCtx.Config.Verify.TurnstileSecret,
				Timeout: 3 * time.Second,
			})
			if verify, err := verifyTurns.Verify(c, req.CfToken, req.IP); err != nil || !verify {
				result.HttpResult(c, nil, errors.Wrapf(xerr.NewErrCode(xerr.TooManyRequests), "verify error: %v", err.Error()))
				return
			}
		}
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := auth.NewUserRegisterLogic(c.Request.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		result.HttpResult(c, resp, err)
	}
}
