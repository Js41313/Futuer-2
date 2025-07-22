package document

import (
	"github.com/Js41313/Futuer-2/internal/logic/app/document"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get document list
func QueryDocumentListHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := document.NewQueryDocumentListLogic(c.Request.Context(), svcCtx)
		resp, err := l.QueryDocumentList()
		result.HttpResult(c, resp, err)
	}
}
