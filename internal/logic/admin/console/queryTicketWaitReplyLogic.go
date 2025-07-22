package console

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type QueryTicketWaitReplyLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewQueryTicketWaitReplyLogic Query ticket wait reply
func NewQueryTicketWaitReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTicketWaitReplyLogic {
	return &QueryTicketWaitReplyLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTicketWaitReplyLogic) QueryTicketWaitReply() (resp *types.TicketWaitRelpyResponse, err error) {
	count, err := l.svcCtx.TicketModel.QueryWaitReplyTotal(l.ctx)
	if err != nil {
		l.Errorw("[QueryTicketWaitReply] Query Database Error: ", logger.Field("error", err.Error()))
		return nil, err
	}
	return &types.TicketWaitRelpyResponse{
		Count: count,
	}, nil
}
