package auth

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type CheckUserTelephoneLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Check user telephone is exist
func NewCheckUserTelephoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserTelephoneLogic {
	return &CheckUserTelephoneLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckUserTelephoneLogic) CheckUserTelephone(req *types.TelephoneCheckUserRequest) (resp *types.TelephoneCheckUserResponse, err error) {
	phoneNumber, err := phone.FormatToE164(req.TelephoneAreaCode, req.Telephone)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TelephoneError), "Invalid phone number")
	}
	authMethods, err := l.svcCtx.UserModel.FindUserAuthMethodByOpenID(l.ctx, "mobile", phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find user by email error: %v", err.Error())
	}

	return &types.TelephoneCheckUserResponse{
		Exist: authMethods.UserId != 0,
	}, nil
}
