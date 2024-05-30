package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	DB               DBConfig      `yaml:"db" env-required:"true"`
	TelegramApiToken string        `yaml:"telegram_api_token" env-required:"true"`
	FetchInterval    time.Duration `yaml:"fetch_interval" env-default:"5h"`
}

type DBConfig struct {
	Host     string `yaml:"db_host" env-required:"true"`
	User     string `yaml:"db_user" env-required:"true"`
	Name     string `yaml:"db_name" env-required:"true"`
	Password string `yaml:"db_password" env-required:"true"`
	Port     string `yaml:"db_port" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("CONFIG FILE does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Cant read CONFIG FILE: " + err.Error())
	}

	return &cfg

}
