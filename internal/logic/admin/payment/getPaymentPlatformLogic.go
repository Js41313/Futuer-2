package payment

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/payment"
)

type GetPaymentPlatformLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get supported payment platform
func NewGetPaymentPlatformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentPlatformLogic {
	return &GetPaymentPlatformLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPaymentPlatformLogic) GetPaymentPlatform() (resp *types.PlatformResponse, err error) {
	resp = &types.PlatformResponse{
		List: payment.GetSupportedPlatforms(),
	}
	return
}
