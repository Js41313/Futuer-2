package tool

import (
	"context"
	"encoding/json"

	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type GetSystemLogLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetSystemLogLogic Get System Log
func NewGetSystemLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemLogLogic {
	return &GetSystemLogLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemLogLogic) GetSystemLog() (resp *types.LogResponse, err error) {
	lines, err := logger.ReadLastNLines(l.svcCtx.Config.Logger.Path, 50)
	if err != nil {
		l.Error(err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR), "get system log error: %v", err.Error())
	}
	var list []map[string]interface{}
	for _, line := range lines {
		var log map[string]interface{}
		if err = json.Unmarshal([]byte(line), &log); err != nil {
			l.Error(err)
			continue
		}
		list = append(list, log)
	}

	return &types.LogResponse{
		List: list,
	}, nil
}
