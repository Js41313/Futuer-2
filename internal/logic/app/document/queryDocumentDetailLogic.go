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

type QueryDocumentDetailLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewQueryDocumentDetailLogic Get document detail
func NewQueryDocumentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDocumentDetailLogic {
	return &QueryDocumentDetailLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDocumentDetailLogic) QueryDocumentDetail(req *types.QueryDocumentDetailRequest) (resp *types.Document, err error) {
	// find document
	data, err := l.svcCtx.DocumentModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Error("[QueryDocumentDetailLogic] FindOne error", logger.Field("id", req.Id), logger.Field("error", err.Error()))
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DatabaseQueryError), "FindOne error: %s", err.Error())
	}
	resp = &types.Document{}
	tool.DeepCopy(resp, data)
	return
}
