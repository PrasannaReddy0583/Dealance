package config

import "os"

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Addr string
}

func Load() *Config {
	return &Config{
		HTTP: HTTPConfig{
			Addr: getEnv("HTTP_ADDR", ":8080"),
		},
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
