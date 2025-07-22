package authMethod

import (
	"context"

	"github.com/Js41313/Futuer-2/pkg/sms"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type GetSmsPlatformLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get sms support platform
func NewGetSmsPlatformLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsPlatformLogic {
	return &GetSmsPlatformLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsPlatformLogic) GetSmsPlatform() (resp *types.PlatformResponse, err error) {
	return &types.PlatformResponse{
		List: sms.GetSupportedPlatforms(),
	}, nil
}
