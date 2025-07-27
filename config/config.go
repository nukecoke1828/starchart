package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port     int    `env:"PORT" envDefault:"8080"`
	RedisURL string `env:"REDIS_URL" envDefault:"redis://localhost:6379"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}