package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type QueryUserInfoLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// query user info
func NewQueryUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserInfoLogic {
	return &QueryUserInfoLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserInfoLogic) QueryUserInfo() (resp *types.UserInfoResponse, err error) {
	u := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	var devices []types.UserDevice
	if len(u.UserDevices) != 0 {
		for _, device := range u.UserDevices {
			devices = append(devices, types.UserDevice{
				Id:         device.Id,
				Identifier: device.Identifier,
				Online:     device.Online,
			})
		}
	}
	var authMeths []types.UserAuthMethod
	authMethods, err := l.svcCtx.UserModel.FindUserAuthMethods(l.ctx, u.Id)
	if err == nil && len(authMeths) != 0 {
		for _, as := range authMethods {
			authMeths = append(authMeths, types.UserAuthMethod{
				AuthType:       as.AuthType,
				AuthIdentifier: as.AuthIdentifier,
			})
		}
	}

	resp = &types.UserInfoResponse{
		Id:          u.Id,
		Balance:     u.Balance,
		Avatar:      u.Avatar,
		ReferCode:   u.ReferCode,
		RefererId:   u.RefererId,
		Devices:     devices,
		AuthMethods: authMeths,
	}
	return
}
