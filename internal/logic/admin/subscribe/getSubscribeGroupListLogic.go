package subscribe

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/subscribe"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetSubscribeGroupListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get subscribe group list
func NewGetSubscribeGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeGroupListLogic {
	return &GetSubscribeGroupListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeGroupListLogic) GetSubscribeGroupList() (resp *types.GetSubscribeGroupListResponse, err error) {
	var list []*subscribe.Group
	var total int64
	err = l.svcCtx.DB.Model(&subscribe.Group{}).Count(&total).Find(&list).Error
	if err != nil {
		l.Logger.Error("[GetSubscribeGroupListLogic] get subscribe group list failed: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe group list failed: %v", err.Error())
	}
	groupList := make([]types.SubscribeGroup, 0)
	tool.DeepCopy(&groupList, list)
	return &types.GetSubscribeGroupListResponse{
		Total: total,
		List:  groupList,
	}, nil
}
