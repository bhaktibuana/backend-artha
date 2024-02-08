package configs

import (
	"os"
)

type SSMTPConfig struct {
	Port     string
	Host     string
	Username string
	Password string
}

func SMTPConfig() SSMTPConfig {
	return SSMTPConfig{
		Port:     os.Getenv("SMTP_PORT"),
		Host:     os.Getenv("SMTP_HOST"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
}
