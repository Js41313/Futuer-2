package document

import (
	"context"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type BatchDeleteDocumentLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Batch delete document
func NewBatchDeleteDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteDocumentLogic {
	return &BatchDeleteDocumentLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteDocumentLogic) BatchDeleteDocument(req *types.BatchDeleteDocumentRequest) error {
	for _, id := range req.Ids {
		if err := l.svcCtx.DocumentModel.Delete(l.ctx, id); err != nil {
			l.Errorw("[BatchDeleteDocument] Database Error", logger.Field("error", err.Error()))
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseDeletedError), "failed to delete document: %v", err.Error())
		}
	}
	return nil
}
