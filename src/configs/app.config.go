package configs

import (
	"os"
)

type IAppConfig struct {
	Port    string
	GinMode string
	BaseUrl string
}

func AppConfig() IAppConfig {
	return IAppConfig{
		Port:    ":" + os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
		BaseUrl: os.Getenv("BASE_URL"),
	}
}
