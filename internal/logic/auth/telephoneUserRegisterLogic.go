package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/jwt"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/uuidx"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CacheKeyPayload struct {
	Code   string `json:"code"`
	LastAt int64  `json:"lastAt"`
}

type TelephoneUserRegisterLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewTelephoneUserRegisterLogic User Telephone register
func NewTelephoneUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TelephoneUserRegisterLogic {
	return &TelephoneUserRegisterLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TelephoneUserRegisterLogic) TelephoneUserRegister(req *types.TelephoneRegisterRequest) (resp *types.LoginResponse, err error) {
	c := l.svcCtx.Config.Register
	// Check if the registration is stopped
	if c.StopRegister {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.StopRegister), "stop register")
	}
	if !phone.Check(req.TelephoneAreaCode, req.Telephone) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TelephoneError), "telephone number error")
	}

	if !l.svcCtx.Config.Mobile.Enable {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SmsNotEnabled), "sms login is not enabled")
	}

	phoneNumber, err := phone.FormatToE164(req.TelephoneAreaCode, req.Telephone)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TelephoneError), "Invalid phone number")
	}

	// if the email verification is enabled, the verification code is required
	cacheKey := fmt.Sprintf("%s:%s:%s", config.AuthCodeTelephoneCacheKey, constant.ParseVerifyType(uint8(constant.Register)), phoneNumber)
	value, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
	if err != nil {
		l.Errorw("Redis Error", logger.Field("error", err.Error()), logger.Field("cacheKey", cacheKey))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}

	var payload CacheKeyPayload
	err = json.Unmarshal([]byte(value), &payload)
	if err != nil {
		l.Errorw("Unmarshal Error", logger.Field("error", err.Error()), logger.Field("value", value))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}
	if payload.Code != req.Code {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.VerifyCodeError), "code error")
	}
	l.svcCtx.Redis.Del(l.ctx, cacheKey)
	// Check if the user exists
	_, err = l.svcCtx.UserModel.FindUserAuthMethodByOpenID(l.ctx, "mobile", phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorw("FindOneByTelephone Error", logger.Field("error", err))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "query user info failed: %v", err.Error())
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserExist), "telephone already exists")
	}
	var referer *user.User
	if req.Invite == "" {
		if l.svcCtx.Config.Invite.ForcedInvite {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.InviteCodeError), "invite code is required")
		}
	} else {
		// Check if the invite code is valid
		referer, err = l.svcCtx.UserModel.FindOneByReferCode(l.ctx, req.Invite)
		if err != nil {
			l.Errorw("FindOneByReferCode Error", logger.Field("error", err))
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.InviteCodeError), "invite code is invalid")
		}
	}

	// Generate password
	pwd := tool.EncodePassWord(req.Password)
	userInfo := &user.User{
		Password: pwd,
		AuthMethods: []user.AuthMethods{
			{
				AuthType:       "mobile",
				AuthIdentifier: phoneNumber,
				Verified:       true,
			},
		},
	}
	if referer != nil {
		userInfo.RefererId = referer.Id
	}
	err = l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		// Save user information
		if err := db.Create(userInfo).Error; err != nil {
			return err
		}
		// Generate ReferCode
		userInfo.ReferCode = uuidx.UserInviteCode(userInfo.Id)
		// Update ReferCode
		if err := db.Model(&user.User{}).Where("id = ?", userInfo.Id).Update("refer_code", userInfo.ReferCode).Error; err != nil {
			return err
		}
		if l.svcCtx.Config.Register.EnableTrial {
			// Active trial
			if err = l.activeTrial(userInfo.Id); err != nil {
				return err
			}
		}
		return nil
	})
	// Generate session id
	sessionId := uuidx.NewUUID().String()
	// Generate token
	token, err := jwt.NewJwtToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.JwtAuth.AccessExpire,
		jwt.WithOption("UserId", userInfo.Id),
		jwt.WithOption("SessionId", sessionId),
	)
	if err != nil {
		l.Logger.Error("[UserLogin] token generate error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "token generate error: %v", err.Error())
	}
	sessionIdCacheKey := fmt.Sprintf("%v:%v", config.SessionIdKey, sessionId)
	if err = l.svcCtx.Redis.Set(l.ctx, sessionIdCacheKey, userInfo.Id, time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)*time.Second).Err(); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "set session id error: %v", err.Error())
	}
	return &types.LoginResponse{
		Token: token,
	}, nil
}

func (l *TelephoneUserRegisterLogic) activeTrial(uid int64) error {
	sub, err := l.svcCtx.SubscribeModel.FindOne(l.ctx, l.svcCtx.Config.Register.TrialSubscribe)
	if err != nil {
		return err
	}
	userSub := &user.Subscribe{
		Id:          0,
		UserId:      uid,
		OrderId:     0,
		SubscribeId: sub.Id,
		StartTime:   time.Now(),
		ExpireTime:  tool.AddTime(l.svcCtx.Config.Register.TrialTimeUnit, l.svcCtx.Config.Register.TrialTime, time.Now()),
		Traffic:     sub.Traffic,
		Download:    0,
		Upload:      0,
		Token:       uuidx.SubscribeToken(fmt.Sprintf("Trial-%v", uid)),
		UUID:        uuidx.NewUUID().String(),
		Status:      1,
	}
	return l.svcCtx.UserModel.InsertSubscribe(l.ctx, userSub)
}
