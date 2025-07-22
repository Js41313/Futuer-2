package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/uuidx"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateUserSubscribeLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create user subcribe
func NewCreateUserSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserSubscribeLogic {
	return &CreateUserSubscribeLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserSubscribeLogic) CreateUserSubscribe(req *types.CreateUserSubscribeRequest) error {
	// validate user
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		l.Errorw("FindOne error", logger.Field("error", err.Error()), logger.Field("userId", req.UserId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOne error: %v", err.Error())
	}
	subs, err := l.svcCtx.UserModel.QueryUserSubscribe(l.ctx, req.UserId)
	if err != nil {
		l.Errorw("QueryUserSubscribe error", logger.Field("error", err.Error()), logger.Field("userId", req.UserId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryUserSubscribe error: %v", err.Error())
	}
	if len(subs) >= 1 && l.svcCtx.Config.Subscribe.SingleModel {
		return errors.Wrapf(xerr.NewErrCode(xerr.SingleSubscribeModeExceedsLimit), "Single subscribe mode exceeds limit")
	}
	sub, err := l.svcCtx.SubscribeModel.FindOne(l.ctx, req.SubscribeId)
	if err != nil {
		l.Errorw("FindOne error", logger.Field("error", err.Error()), logger.Field("subscribeId", req.SubscribeId))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOne error: %v", err.Error())
	}
	if req.Traffic == 0 {
		req.Traffic = sub.Traffic
	}

	userSub := user.Subscribe{
		UserId:      req.UserId,
		SubscribeId: req.SubscribeId,
		StartTime:   time.Now(),
		ExpireTime:  time.UnixMilli(req.ExpiredAt),
		Traffic:     req.Traffic,
		Download:    0,
		Upload:      0,
		Token:       uuidx.SubscribeToken(fmt.Sprintf("adminCreate:%d", time.Now().UnixMilli())),
		UUID:        uuid.New().String(),
		Status:      1,
	}
	if err = l.svcCtx.UserModel.InsertSubscribe(l.ctx, &userSub); err != nil {
		l.Errorw("InsertSubscribe error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "InsertSubscribe error: %v", err.Error())
	}

	err = l.svcCtx.UserModel.UpdateUserCache(l.ctx, userInfo)
	if err != nil {
		l.Errorw("UpdateUserCache error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "UpdateUserCache error: %v", err.Error())
	}

	return nil
}
