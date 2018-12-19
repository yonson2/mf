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

var Env = getEnv("ENV", "development")
var ApiUrl = getEnv("LF_API_URL", "https://leetapi.pramos.me")
var HttpPort = getEnv("LF_HTTP_SERVER_PORT", "9120")
