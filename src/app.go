package app

import (
	"api-artha/src/models"
	"api-artha/src/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DBConnection() {
	models.ConnectDatabase()
}

func Middlewares(app *gin.Engine) {
	// Middleware to set the trusted headers (trust proxy)
	app.Use(func(context *gin.Context) {
		context.Request.Header.Set("X-Real-IP", context.GetHeader("X-Real-IP"))
		context.Request.Header.Set("X-Forwarded-For", context.GetHeader("X-Forwarded-For"))
		context.Request.Header.Set("X-Forwarded-Proto", context.GetHeader("X-Forwarded-Proto"))
		context.Next()
	})

	// Middleware to disable Cross-Origin Embedder Policy
	app.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Cross-Origin-Embedder-Policy", "unsafe-none")
		context.Next()
	})

	app.Use(cors.Default())
}

func Routes(app *gin.Engine) {
	routers.Index(app)
}

func ListenServer(app *gin.Engine, port string) {
	app.Run(port)
}
