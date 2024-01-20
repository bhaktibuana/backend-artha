package main

import (
	app "api-artha/src"
	"api-artha/src/configs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.SetMode(configs.AppConfig().GinMode)
	server := gin.Default()

	app.DBConnection()
	app.Middlewares(server)
	app.Routes(server)
	app.ListenServer(server, configs.AppConfig().Port)
}
