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

type GetDocumentDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get document detail
func NewGetDocumentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentDetailLogic {
	return &GetDocumentDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentDetailLogic) GetDocumentDetail(req *types.GetDocumentDetailRequest) (resp *types.Document, err error) {
	data, err := l.svcCtx.DocumentModel.QueryDocumentDetail(l.ctx, req.Id)
	if err != nil {
		l.Errorw("[GetDocumentDetail] Database Error", logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "QueryDocumentDetail error: %v", err.Error())
	}
	resp = &types.Document{
		Id:        data.Id,
		Title:     data.Title,
		Tags:      tool.StringMergeAndRemoveDuplicates(data.Tags),
		Content:   data.Content,
		CreatedAt: data.CreatedAt.UnixMilli(),
		UpdatedAt: data.UpdatedAt.UnixMilli(),
	}
	return
}
