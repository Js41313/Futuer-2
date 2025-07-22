package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/perfect-panel/server/pkg/logger"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/model/user"
	"github.com/perfect-panel/server/pkg/jwt"
	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/uuidx"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
)

type UserLoginLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserLoginLogic User login
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.LoginResponse, err error) {
	loginStatus := false
	var userInfo *user.User
	// Record login status
	defer func(svcCtx *svc.ServiceContext) {
		if userInfo.Id != 0 {
			if err := svcCtx.UserModel.InsertLoginLog(l.ctx, &user.LoginLog{
				UserId:    userInfo.Id,
				LoginIP:   req.IP,
				UserAgent: req.UserAgent,
				Success:   &loginStatus,
			}); err != nil {
				l.Logger.Error("[UserLogin] insert login log error", logger.Field("error", err.Error()))
			}
		}
	}(l.svcCtx)

	userInfo, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			logger.WithContext(l.ctx).Error(err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotExist), "user email not exist: %v", req.Email)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "query user info failed: %v", err.Error())
	}
	// Verify password
	if !tool.VerifyPassWord(req.Password, userInfo.Password) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserPasswordError), "user password")
	}
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
	loginStatus = true
	return &types.LoginResponse{
		Token: token,
	}, nil
}
