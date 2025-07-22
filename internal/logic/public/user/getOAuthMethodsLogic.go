package user

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetOAuthMethodsLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get OAuth Methods
func NewGetOAuthMethodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAuthMethodsLogic {
	return &GetOAuthMethodsLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOAuthMethodsLogic) GetOAuthMethods() (resp *types.GetOAuthMethodsResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		logger.Error("current user is not found in context")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	methods, err := l.svcCtx.UserModel.FindUserAuthMethods(l.ctx, u.Id)
	if err != nil {
		l.Errorw("find user auth methods failed:", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "find user auth methods failed: %v", err.Error())
	}
	list := make([]types.UserAuthMethod, 0)
	tool.DeepCopy(&list, methods)
	return &types.GetOAuthMethodsResponse{
		Methods: list,
	}, nil
}
