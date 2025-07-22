package log

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/log"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetMessageLogListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetMessageLogListLogic Get message log list
func NewGetMessageLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogListLogic {
	return &GetMessageLogListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageLogListLogic) GetMessageLogList(req *types.GetMessageLogListRequest) (resp *types.GetMessageLogListResponse, err error) {
	total, data, err := l.svcCtx.LogModel.FindMessageLogList(l.ctx, req.Page, req.Size, log.MessageLogFilterParams{
		Type:     req.Type,
		Platform: req.Platform,
		To:       req.To,
		Subject:  req.Subject,
		Content:  req.Content,
		Status:   req.Status,
	})
	if err != nil {
		l.Errorw("[GetMessageLogList] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "[GetMessageLogList] Database Error: %s", err.Error())
	}
	var list []types.MessageLog
	tool.DeepCopy(&list, data)

	return &types.GetMessageLogListResponse{
		Total: total,
		List:  list,
	}, nil
}
