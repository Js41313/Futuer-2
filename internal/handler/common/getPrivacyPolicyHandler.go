package common

import (
	"github.com/Js41313/Futuer-2/internal/logic/common"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/result"
	"github.com/gin-gonic/gin"
)

// Get Privacy Policy
func GetPrivacyPolicyHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		l := common.NewGetPrivacyPolicyLogic(c.Request.Context(), svcCtx)
		resp, err := l.GetPrivacyPolicy()
		result.HttpResult(c, resp, err)
	}
}
