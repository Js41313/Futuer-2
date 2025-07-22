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

type GetUserSubscribeByIdLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user subcribe by id
func NewGetUserSubscribeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSubscribeByIdLogic {
	return &GetUserSubscribeByIdLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSubscribeByIdLogic) GetUserSubscribeById(req *types.GetUserSubscribeByIdRequest) (resp *types.UserSubscribeDetail, err error) {
	sub, err := l.svcCtx.UserModel.FindOneSubscribeDetailsById(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[GetUserSubscribeByIdLogic] FindOneSubscribeDetailsById error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOneSubscribeDetailsById error: %v", err.Error())
	}
	var subscribeDetails types.UserSubscribeDetail
	tool.DeepCopy(&subscribeDetails, sub)
	return &subscribeDetails, nil
}
