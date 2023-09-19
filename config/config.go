package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	APIServer
	MongoDB
	TokenPair
}

type APIServer struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`
	BindAddr string `env:"BIND_ADDR" env-default:":8080"`
}

type MongoDB struct {
	Host       string `env:"HOST" env-default:"host.docker.internal"`
	Port       string `env:"PORT" env-default:"27017"`
	User       string `env:"USER" env-default:"mongodb"`
	Password   string `env:"PASSWORD" env-default:"mongodb"`
	DBName     string `env:"DB_NAME" env-default:"mongodb"`
	Collection string `env:"COLLECTION" env-default:"user_authentication"`
}

type TokenPair struct {
	AccessSecretKey string `env:"ACCESS_SECRET_KEY" env-default:"key123"`

	//token lifetime in minutes
	AccessTokenLifetime  int `env:"ACCESS_TOKEN_LIFETIME" env-default:"30"`      // 30 minutes
	RefreshTokenLifetime int `env:"REFRESH_TOKEN_LIFETIME" env-default:"525600"` // 365 days
}

func GetConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read the config: %s", err)
	}
	return &cfg
}
