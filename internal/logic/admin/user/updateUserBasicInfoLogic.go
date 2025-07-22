package user

import (
	"context"
	"os"
	"strings"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateUserBasicInfoLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateUserBasicInfoLogic Update user basic info
func NewUpdateUserBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBasicInfoLogic {
	return &UpdateUserBasicInfoLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBasicInfoLogic) UpdateUserBasicInfo(req *types.UpdateUserBasiceInfoRequest) error {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		l.Errorw("[UpdateUserBasicInfoLogic] Find User Error:", logger.Field("err", err.Error()), logger.Field("userId", req.UserId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Find User Error")
	}

	isDemo := strings.ToLower(os.Getenv("PPANEL_MODE")) == "demo"

	tool.DeepCopy(userInfo, req)
	if req.Avatar != "" && !tool.IsValidImageSize(req.Avatar, 1024) {
		return errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "Invalid Image Size")
	}
	userInfo.Balance = req.Balance
	userInfo.GiftAmount = req.GiftAmount
	userInfo.Commission = req.Commission

	if req.Password != "" {
		if userInfo.Id == 2 && isDemo {
			return errors.Wrapf(xerr.NewErrCodeMsg(503, "Demo mode does not allow modification of the admin user password"), "UpdateUserBasicInfo failed: cannot update admin user password in demo mode")
		}
		userInfo.Password = tool.EncodePassWord(req.Password)
	}

	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		l.Errorw("[UpdateUserBasicInfoLogic] Update User Error:", logger.Field("err", err.Error()), logger.Field("userId", req.UserId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "Update User Error")
	}

	return nil
}
