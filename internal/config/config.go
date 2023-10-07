package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	defaultServerAddress = ":8080"
	defaultSQLitePath    = ":memory:"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Auth   AuthConfig   `mapstructure:"auth"`
	Env    EnvConfig
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
}

type AuthConfig struct {
	AccessTokenLifetime  time.Duration `mapstructure:"access_token_lifetime"`
	RefreshTokenLifetime time.Duration `mapstructure:"refresh_token_lifetime"`
}

type EnvConfig struct {
	SQLitePath    string `envconfig:"SQLITE_PATH"`
	JwtSigningKey string `envconfig:"JWT_SIGNING_KEY"`
}

func Load(path, name string) (*Config, error) {
	viper.SetDefault("server.address", defaultServerAddress)
	viper.SetDefault("sqlite.path", defaultSQLitePath)

	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	godotenv.Load()
	if err := envconfig.Process("", &cfg.Env); err != nil {
		return nil, err
	}

	return cfg, nil
}
