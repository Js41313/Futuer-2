package ticket

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"github.com/Js41313/Futuer-2/pkg/logger"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetUserTicketListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get ticket list
func NewGetUserTicketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTicketListLogic {
	return &GetUserTicketListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserTicketListLogic) GetUserTicketList(req *types.GetUserTicketListRequest) (resp *types.GetUserTicketListResponse, err error) {
	u, ok := l.ctx.Value(constant.CtxKeyUser).(*user.User)
	if !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidAccess), "Invalid Access")
	}
	l.Logger.Debugf("Current user: %v", u.Id)
	total, list, err := l.svcCtx.TicketModel.QueryTicketList(l.ctx, req.Page, req.Size, u.Id, req.Status, req.Search)
	if err != nil {
		l.Errorw("[GetUserTicketListLogic] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryTicketList error: %v", err)
	}
	resp = &types.GetUserTicketListResponse{
		Total: total,
		List:  make([]types.Ticket, 0),
	}
	tool.DeepCopy(&resp.List, list)
	return
}
