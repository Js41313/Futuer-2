package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/pkg/constant"
	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type UpdateBindMobileLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update Bind Mobile
func NewUpdateBindMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBindMobileLogic {
	return &UpdateBindMobileLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBindMobileLogic) UpdateBindMobile(req *types.UpdateBindMobileRequest) error {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	// verify mobile
	phoneNumber, err := phone.FormatToE164(req.AreaCode, req.Mobile)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.TelephoneError), "Invalid phone number")
	}
	cacheKey := fmt.Sprintf("%s:%s:%s", config.AuthCodeTelephoneCacheKey, constant.Register, phoneNumber)
	code, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
	if err != nil {
		l.Errorw("Redis Error", logger.Field("error", err.Error()), logger.Field("cacheKey", cacheKey))
		return errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}
	var payload CacheKeyPayload
	err = json.Unmarshal([]byte(code), &payload)
	if err != nil {
		l.Errorw("Redis Error", logger.Field("error", err.Error()), logger.Field("cacheKey", cacheKey))
		return errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}
	if payload.Code != req.Code {
		return errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}
	l.svcCtx.Redis.Del(l.ctx, cacheKey)

	m, err := l.svcCtx.UserModel.FindUserAuthMethodByOpenID(l.ctx, "mobile", req.Mobile)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindUserAuthMethodByOpenID error")
	}
	if m.Id > 0 {
		return errors.Wrapf(xerr.NewErrCode(xerr.UserExist), "mobile already bind")
	}

	method, err := l.svcCtx.UserModel.FindUserAuthMethodByUserId(l.ctx, "mobile", u.Id)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindUserAuthMethodByOpenID error")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		method = &user.AuthMethods{
			UserId:         u.Id,
			AuthType:       "mobile",
			AuthIdentifier: req.Mobile,
			Verified:       true,
		}
		if err := l.svcCtx.UserModel.InsertUserAuthMethods(l.ctx, method); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "InsertUserAuthMethods error")
		}
	} else {
		method.Verified = true
		method.AuthIdentifier = req.Mobile
		if err := l.svcCtx.UserModel.UpdateUserAuthMethods(l.ctx, method); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "UpdateUserAuthMethods error")
		}
	}
	return nil
}
