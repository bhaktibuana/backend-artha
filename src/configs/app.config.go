package configs

import (
	"os"
)

type SAppConfig struct {
	Port    string
	GinMode string
	BaseUrl string
}

func AppConfig() SAppConfig {
	return SAppConfig{
		Port:    ":" + os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
		BaseUrl: os.Getenv("BASE_URL"),
	}
}
