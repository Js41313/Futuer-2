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

type GetAnnouncementListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get announcement list
func NewGetAnnouncementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnouncementListLogic {
	return &GetAnnouncementListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnnouncementListLogic) GetAnnouncementList(req *types.GetAnnouncementListRequest) (resp *types.GetAnnouncementListResponse, err error) {
	total, list, err := l.svcCtx.AnnouncementModel.GetAnnouncementListByPage(l.ctx, int(req.Page), int(req.Size), announcement.Filter{
		Show:   req.Show,
		Pinned: req.Pinned,
		Popup:  req.Popup,
		Search: req.Search,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "GetAnnouncementListByPage error: %v", err.Error())
	}
	resp = &types.GetAnnouncementListResponse{}
	resp.Total = total
	resp.List = make([]types.Announcement, 0)
	tool.DeepCopy(&resp.List, list)
	return
}
