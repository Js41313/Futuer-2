package user

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetUserAuthMethodLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user auth method
func NewGetUserAuthMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthMethodLogic {
	return &GetUserAuthMethodLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAuthMethodLogic) GetUserAuthMethod(req *types.GetUserAuthMethodRequest) (resp *types.GetUserAuthMethodResponse, err error) {
	methods, err := l.svcCtx.UserModel.FindUserAuthMethods(l.ctx, req.UserId)
	if err != nil {
		l.Errorw("[GetUserAuthMethodLogic] Get User Auth Method Error:", logger.Field("err", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "Get User Auth Method Error")
	}
	list := make([]types.UserAuthMethod, 0)
	tool.DeepCopy(&list, methods)

	return &types.GetUserAuthMethodResponse{
		AuthMethods: list,
	}, nil
}
