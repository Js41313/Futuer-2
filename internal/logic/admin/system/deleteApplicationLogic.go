package system

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteApplicationLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApplicationLogic {
	return &DeleteApplicationLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApplicationLogic) DeleteApplication(req *types.DeleteApplicationRequest) error {
	// delete application
	err := l.svcCtx.ApplicationModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[DeleteApplicationLogic] delete application error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete application error: %v", err.Error())
	}
	return nil
}
