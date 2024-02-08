package configs

import (
	"os"
)

type SClientConfig struct {
	ArthaUrl string
}

func ClientConfig() SClientConfig {
	return SClientConfig{
		ArthaUrl: os.Getenv("ARTHA_URL"),
	}
}
