package announcement

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/model/announcement"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type QueryAnnouncementLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewQueryAnnouncementLogic Query announcement
func NewQueryAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAnnouncementLogic {
	return &QueryAnnouncementLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAnnouncementLogic) QueryAnnouncement(req *types.QueryAnnouncementRequest) (resp *types.QueryAnnouncementResponse, err error) {
	enable := true
	total, list, err := l.svcCtx.AnnouncementModel.GetAnnouncementListByPage(l.ctx, req.Page, req.Size, announcement.Filter{
		Show:   &enable,
		Pinned: req.Pinned,
		Popup:  req.Popup,
	})
	if err != nil {
		l.Error("[QueryAnnouncementLogic] GetAnnouncementListByPage error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetAnnouncementListByPage error: %v", err.Error())
	}
	resp = &types.QueryAnnouncementResponse{}
	resp.Total = total
	resp.List = make([]types.Announcement, 0)
	tool.DeepCopy(&resp.List, list)
	return
}
