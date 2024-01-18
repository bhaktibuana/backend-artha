package routers

import (
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

		context.Set("scheme", scheme)
		context.Next()
	})

	appRouter.AppRouters("/api", router)

	router.NoRoute(func(context *gin.Context) {
		scheme, _ := context.Get("scheme")
		url := fmt.Sprintf("%s://%s%s", scheme, context.Request.Host, context.Request.URL.Path)
		helpers.Response("URL not found", http.StatusNotFound, context, map[string]interface{}{"url": url})
	})

	router.GET("/", func(context *gin.Context) {
		scheme, _ := context.Get("scheme")
		url := fmt.Sprintf("%s://%s", scheme, context.Request.Host)
		helpers.Response("Artha API Service", http.StatusOK, context, map[string]interface{}{"url": url})
	})
}
