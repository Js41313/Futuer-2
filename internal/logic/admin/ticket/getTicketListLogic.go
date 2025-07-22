package ticket

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetTicketListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get ticket list
func NewGetTicketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTicketListLogic {
	return &GetTicketListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTicketListLogic) GetTicketList(req *types.GetTicketListRequest) (resp *types.GetTicketListResponse, err error) {
	total, list, err := l.svcCtx.TicketModel.QueryTicketList(l.ctx, int(req.Page), int(req.Size), req.UserId, req.Status, req.Search)
	if err != nil {
		l.Errorw("[GetTicketList] Query Database Error: ", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryTicketList error: %v", err)
	}
	resp = &types.GetTicketListResponse{
		Total: total,
		List:  make([]types.Ticket, 0),
	}
	tool.DeepCopy(&resp.List, list)
	return
}
