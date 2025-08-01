package common

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/authmethod"
	"github.com/Js41313/Futuer-2/pkg/constant"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CheckVerificationCodeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Check verification code
func NewCheckVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVerificationCodeLogic {
	return &CheckVerificationCodeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckVerificationCodeLogic) CheckVerificationCode(req *types.CheckVerificationCodeRequest) (resp *types.CheckVerificationCodeRespone, err error) {
	resp = &types.CheckVerificationCodeRespone{}
	if req.Method == authmethod.Email {
		cacheKey := fmt.Sprintf("%s:%s:%s", config.AuthCodeCacheKey, constant.ParseVerifyType(req.Type), req.Account)
		value, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
		if err != nil {
			return resp, nil
		}
		var payload CacheKeyPayload
		if err := json.Unmarshal([]byte(value), &payload); err != nil {
			return resp, nil
		}
		if payload.Code != req.Code {
			return resp, nil
		}
		resp.Status = true
	}
	if req.Method == authmethod.Mobile {
		if !phone.CheckPhone(req.Account) {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.TelephoneError), "Invalid phone number")
		}
		cacheKey := fmt.Sprintf("%s:%s:+%s", config.AuthCodeTelephoneCacheKey, constant.ParseVerifyType(req.Type), req.Account)
		value, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
		if err != nil {
			return resp, nil
		}
		var payload CacheKeyPayload
		if err := json.Unmarshal([]byte(value), &payload); err != nil {
			return resp, nil
		}
		if payload.Code != req.Code {
			return resp, nil
		}
		resp.Status = true
	}
	return resp, nil
}
