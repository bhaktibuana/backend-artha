package configs

import (
	"os"
)

type SAppConfig struct {
	Port      string
	GinMode   string
	BaseUrl   string
	JwtSecret string
}

func AppConfig() SAppConfig {
	return SAppConfig{
		Port:      ":" + os.Getenv("PORT"),
		GinMode:   os.Getenv("GIN_MODE"),
		BaseUrl:   os.Getenv("BASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET_KEY"),
	}
}
