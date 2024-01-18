package configs

import (
	"os"
)

type IAppConfig struct {
	Port    string
	GinMode string
}

func AppConfig() IAppConfig {
	return IAppConfig{
		Port:    ":" + os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
	}
}
