package config

import (
	"github.com/caarlos0/env/v11"
)

type DbConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`
}

type Config struct {
	Db        DbConfig `envPrefix:""`
	HTTPPort  string   `env:"HTTP_PORT"`
	RedisAddr string   `env:"REDIS_ADDR"`
}

func ReadEnv() (*Config, error) {
	cfg := &Config{}
	opts := env.Options{
		RequiredIfNoDef: true,
	}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}
