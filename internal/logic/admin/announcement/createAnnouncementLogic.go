package announcement

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/announcement"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateAnnouncementLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create announcement
func NewCreateAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAnnouncementLogic {
	return &CreateAnnouncementLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAnnouncementLogic) CreateAnnouncement(req *types.CreateAnnouncementRequest) error {

	if err := l.svcCtx.AnnouncementModel.Insert(l.ctx, &announcement.Announcement{
		Title:   req.Title,
		Content: req.Content,
	}); err != nil {
		l.Errorw("[CreateAnnouncement] Database Error", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create announcement failed: %v", err.Error())
	}

	return nil
}
