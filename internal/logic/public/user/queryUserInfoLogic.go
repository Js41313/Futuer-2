package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/tool"
)

type QueryUserInfoLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Query User Info
func NewQueryUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserInfoLogic {
	return &QueryUserInfoLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserInfoLogic) QueryUserInfo() (resp *types.User, err error) {
	resp = &types.User{}
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	tool.DeepCopy(resp, u)

	var userMethods []types.UserAuthMethod
	for _, method := range resp.AuthMethods {
		var item types.UserAuthMethod
		tool.DeepCopy(&item, method)

		switch method.AuthType {
		case "mobile":
			item.AuthIdentifier = phone.MaskPhoneNumber(method.AuthIdentifier)
		case "email":
		default:
			item.AuthIdentifier = maskOpenID(method.AuthIdentifier)
		}
		userMethods = append(userMethods, item)
	}
	resp.AuthMethods = userMethods
	return resp, nil
}

// maskOpenID 脱敏 OpenID，只保留前 3 和后 3 位
func maskOpenID(openID string) string {
	length := len(openID)
	if length <= 6 {
		return "***" // 如果 ID 太短，直接返回 "***"
	}

	// 计算中间需要被替换的 `*` 数量
	maskLength := length - 6
	mask := make([]byte, maskLength)
	for i := range mask {
		mask[i] = '*'
	}

	// 组合脱敏后的 OpenID
	return openID[:3] + string(mask) + openID[length-3:]
}
