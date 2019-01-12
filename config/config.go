package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var HttpPort = getEnv("MF_HTTP_SERVER_PORT", "9120")
