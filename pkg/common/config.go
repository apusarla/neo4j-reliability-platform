package common

import (
	"log"
	"os"
)

type Config struct {
	ServiceName string
	HTTPPort    string
	Env         string
}

func LoadConfig(defaultServiceName, defaultPort string) Config {
	cfg := Config{
		ServiceName: getEnv("SERVICE_NAME", defaultServiceName),
		HTTPPort:    getEnv("HTTP_PORT", defaultPort),
		Env:         getEnv("ENV", "dev"),
	}

	log.Printf("Loaded config: service=%s port=%s env=%s",
		cfg.ServiceName, cfg.HTTPPort, cfg.Env)

	return cfg
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
