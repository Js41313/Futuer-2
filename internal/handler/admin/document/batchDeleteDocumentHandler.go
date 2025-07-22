package document

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/document"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Batch delete document
func BatchDeleteDocumentHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.BatchDeleteDocumentRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := document.NewBatchDeleteDocumentLogic(c.Request.Context(), svcCtx)
		err := l.BatchDeleteDocument(&req)
		result.HttpResult(c, nil, err)
	}
}
