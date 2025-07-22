package authMethod

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/email"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type GetEmailPlatformLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get email support platform
func NewGetEmailPlatformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailPlatformLogic {
	return &GetEmailPlatformLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailPlatformLogic) GetEmailPlatform() (resp *types.PlatformResponse, err error) {
	return &types.PlatformResponse{
		List: email.GetSupportedPlatforms(),
	}, nil
}
