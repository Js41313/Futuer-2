package common

import (
	"github.com/Js41313/Futuer-2/internal/logic/common"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Tos Content
func GetApplicationHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := common.NewGetApplicationLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetApplication()
		result.HttpResult(c, resp, err)
	}
}
