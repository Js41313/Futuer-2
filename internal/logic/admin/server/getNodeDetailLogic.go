package server

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

type GetNodeDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeDetailLogic {
	return &GetNodeDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeDetailLogic) GetNodeDetail(req *types.GetDetailRequest) (resp *types.Server, err error) {
	detail, err := l.svcCtx.ServerModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "get server detail error: %v", err.Error())
	}
	resp = &types.Server{}
	tool.DeepCopy(resp, detail)
	var cfg map[string]interface{}
	err = json.Unmarshal([]byte(detail.Config), &cfg)
	if err != nil {
		cfg = make(map[string]interface{})
	}
	resp.Config = cfg
	return
}
