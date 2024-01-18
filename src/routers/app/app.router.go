package appRouter

import "github.com/gin-gonic/gin"

func AppRouters(path string, router *gin.Engine) {
	router.Group(path)
	{

	}
}
