package user

import (
	"context"

	"github.com/perfect-panel/server/pkg/tool"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/logger"
)

type GetUserSubscribeDevicesLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user subcribe devices
func NewGetUserSubscribeDevicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSubscribeDevicesLogic {
	return &GetUserSubscribeDevicesLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSubscribeDevicesLogic) GetUserSubscribeDevices(req *types.GetUserSubscribeDevicesRequest) (resp *types.GetUserSubscribeDevicesResponse, err error) {
	list, total, err := l.svcCtx.UserModel.QueryDevicePageList(l.ctx, req.UserId, req.SubscribeId, req.Page, req.Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetUserSubscribeDevices failed: %v", err.Error())
	}
	userRespList := make([]types.UserDevice, 0)
	tool.DeepCopy(&userRespList, list)
	return &types.GetUserSubscribeDevicesResponse{
		Total: total,
		List:  userRespList,
	}, nil
}
