package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type BindTelegramLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Bind Telegram
func NewBindTelegramLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindTelegramLogic {
	return &BindTelegramLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindTelegramLogic) BindTelegram() (resp *types.BindTelegramResponse, err error) {
	session := l.ctx.Value("session").(string)
	return &types.BindTelegramResponse{
		Url:       fmt.Sprintf("https://t.me/%s?start=%s", l.svcCtx.Config.Telegram.BotName, session),
		ExpiredAt: time.Now().Add(300 * time.Second).UnixMilli(),
	}, nil
}
