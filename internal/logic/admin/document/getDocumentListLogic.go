package document

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type GetDocumentListLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get document list
func NewGetDocumentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentListLogic {
	return &GetDocumentListLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentListLogic) GetDocumentList(req *types.GetDocumentListRequest) (resp *types.GetDocumentListResponse, err error) {
	total, data, err := l.svcCtx.DocumentModel.QueryDocumentList(l.ctx, int(req.Page), int(req.Size), req.Tag, req.Search)
	if err != nil {
		l.Errorw("[GetDocumentList] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryDocumentList error: %v", err.Error())
	}
	resp = &types.GetDocumentListResponse{
		Total: total,
		List:  make([]types.Document, 0),
	}
	for _, v := range data {
		resp.List = append(resp.List, types.Document{
			Id:        v.Id,
			Title:     v.Title,
			Tags:      tool.StringMergeAndRemoveDuplicates(v.Tags),
			Content:   v.Content,
			Show:      *v.Show,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
		})
	}
	return
}
