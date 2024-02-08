package routers

import (
	"api-artha/src/configs"
	"api-artha/src/helpers"
	appRouter "api-artha/src/routers/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(router *gin.Engine) {
	router.Use(func(context *gin.Context) {
		scheme := context.Request.Header.Get("X-Forwarded-Proto")

		if scheme == "" {
			scheme = "http"
		}

		if production := configs.AppConfig().GinMode == "release"; production == true {
			baseUrl := configs.AppConfig().BaseUrl
			context.Set("baseUrl", baseUrl)
		} else {
			baseUrl := fmt.Sprintf("%s://%s", scheme, context.Request.Host)
			context.Set("baseUrl", baseUrl)
		}

		context.Next()
	})

	appRouter.AppRouters("/api", router)

	router.NoRoute(func(context *gin.Context) {
		baseUrl, _ := context.Get("baseUrl")
		url := fmt.Sprintf("%s%s", baseUrl, context.Request.URL.Path)
		helpers.Response("URL not found", http.StatusNotFound, context, map[string]interface{}{"url": url})
	})

	router.GET("/", func(context *gin.Context) {
		baseUrl, _ := context.Get("baseUrl")
		url := fmt.Sprintf("%s", baseUrl)
		helpers.Response("Artha API Service", http.StatusOK, context, map[string]interface{}{"url": url})
	})

}
