package subscribe

import (
	"context"
	"encoding/json"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetSubscribeDetailsLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get subscribe details
func NewGetSubscribeDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeDetailsLogic {
	return &GetSubscribeDetailsLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeDetailsLogic) GetSubscribeDetails(req *types.GetSubscribeDetailsRequest) (resp *types.Subscribe, err error) {
	sub, err := l.svcCtx.SubscribeModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Error("[GetSubscribeDetailsLogic] get subscribe details failed: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get subscribe details failed: %v", err.Error())
	}
	resp = &types.Subscribe{}
	tool.DeepCopy(resp, sub)
	if sub.Discount != "" {
		err = json.Unmarshal([]byte(sub.Discount), &resp.Discount)
		if err != nil {
			l.Logger.Error("[GetSubscribeDetailsLogic] JSON unmarshal failed: ", logger.Field("error", err.Error()), logger.Field("discount", sub.Discount))
		}
	}
	resp.Server = tool.StringToInt64Slice(sub.Server)
	resp.ServerGroup = tool.StringToInt64Slice(sub.ServerGroup)
	return resp, nil
}
