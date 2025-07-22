package document

import (
	"github.com/Js41313/Futuer-2/internal/logic/admin/document"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get document detail
func GetDocumentDetailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.GetDocumentDetailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := document.NewGetDocumentDetailLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetDocumentDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
