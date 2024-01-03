package apiserver

import "web-server/internal/app/store"

type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Addr:     "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
