package configs

import (
	"os"
)

type SDbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

func DbConfig() SDbConfig {
	return SDbConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}
