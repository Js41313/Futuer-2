package auth

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type CheckLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Check Account
func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.AppAuthCheckRequest) (resp *types.AppAuthCheckResponse, err error) {
	resp = &types.AppAuthCheckResponse{}
	_, err = findUserByMethod(l.ctx, l.svcCtx, req.Method, req.Identifier, req.Account, req.AreaCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Status = false
			return resp, nil
		}
		return resp, err
	}
	resp.Status = true
	return
}
