package system

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type DeleteApplicationVersionLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete application
func NewDeleteApplicationVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApplicationVersionLogic {
	return &DeleteApplicationVersionLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApplicationVersionLogic) DeleteApplicationVersion(req *types.DeleteApplicationVersionRequest) error {
	// delete application
	err := l.svcCtx.ApplicationModel.DeleteVersion(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[DeleteApplicationVersion] delete application version error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "delete application version error: %v", err.Error())
	}
	return nil
}
