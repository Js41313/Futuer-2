package user

import (
	"context"
	"os"
	"strings"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type BatchDeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger.Logger
}

func NewBatchDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteUserLogic {
	return &BatchDeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.WithContext(ctx),
	}
}

func (l *BatchDeleteUserLogic) BatchDeleteUser(req *types.BatchDeleteUserRequest) error {
	isDemo := strings.ToLower(os.Getenv("PPANEL_MODE")) == "demo"

	if tool.Contains(req.Ids, 2) && isDemo {
		return errors.Wrapf(xerr.NewErrCodeMsg(503, "Demo mode does not allow deletion of the admin user"), "BatchDeleteUser failed: cannot delete admin user in demo mode")
	}

	err := l.svcCtx.UserModel.BatchDeleteUser(l.ctx, req.Ids)
	if err != nil {
		l.Logger.Error("[BatchDeleteUserLogic] BatchDeleteUser failed: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "BatchDeleteUser failed: %v", err.Error())
	}
	return nil
}
