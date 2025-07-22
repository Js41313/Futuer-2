package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/phone"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.WithContext(ctx),
	}
}
func (l *GetUserListLogic) GetUserList(req *types.GetUserListRequest) (*types.GetUserListResponse, error) {
	list, total, err := l.svcCtx.UserModel.QueryPageList(l.ctx, req.Page, req.Size, &user.UserFilterParams{
		UserId:          req.UserId,
		Search:          req.Search,
		SubscribeId:     req.SubscribeId,
		UserSubscribeId: req.UserSubscribeId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetUserListLogic failed: %v", err.Error())
	}

	userRespList := make([]types.User, 0, len(list))

	for _, item := range list {
		var user types.User
		tool.DeepCopy(&user, item)

		// 处理 AuthMethods
		authMethods := make([]types.UserAuthMethod, len(user.AuthMethods)) // 直接创建目标 slice
		for i, method := range user.AuthMethods {
			tool.DeepCopy(&authMethods[i], method)
			if method.AuthType == "mobile" {
				authMethods[i].AuthIdentifier = phone.FormatToInternational(method.AuthIdentifier)
			}
		}
		user.AuthMethods = authMethods

		userRespList = append(userRespList, user)
	}

	return &types.GetUserListResponse{
		Total: total,
		List:  userRespList,
	}, nil
}
