package appRouter

import (
	authRouter "api-artha/src/routers/auth"

	"github.com/gin-gonic/gin"
)

func AppRouters(path string, router *gin.Engine) {
	apiGroup := router.Group(path)
	{
		authRouter.Routes("/auth", apiGroup)
	}
}
