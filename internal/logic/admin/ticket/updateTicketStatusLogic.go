package ticket

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type UpdateTicketStatusLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update ticket status
func NewUpdateTicketStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTicketStatusLogic {
	return &UpdateTicketStatusLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTicketStatusLogic) UpdateTicketStatus(req *types.UpdateTicketStatusRequest) error {

	err := l.svcCtx.TicketModel.UpdateTicketStatus(l.ctx, req.Id, 0, *req.Status)
	if err != nil {
		l.Errorw("[UpdateTicketStatus] Update Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "update ticket error: %v", err.Error())
	}
	return nil
}
