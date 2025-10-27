package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Address    string `env:"SERVER_ADDRESS" envDefault:"localhost:3000"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"myuser"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"mypassword"`
	DBName     string `env:"DB_NAME" envDefault:"mydb"`
	RedisHost  string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort  string `env:"REDIS_PORT" envDefault:"6379"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
