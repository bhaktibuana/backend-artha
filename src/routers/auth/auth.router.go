package authRouter

import (
	authController "api-artha/src/controllers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(basePath string, router *gin.RouterGroup) {
	authGroup := router.Group(basePath)
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}
}
