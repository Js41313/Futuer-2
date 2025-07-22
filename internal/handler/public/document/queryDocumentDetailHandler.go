package document

import (
	"github.com/Js41313/Futuer-2/internal/logic/public/document"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get document detail
func QueryDocumentDetailHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.QueryDocumentDetailRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := document.NewQueryDocumentDetailLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryDocumentDetail(&req)
		result.HttpResult(c, resp, err)
	}
}
