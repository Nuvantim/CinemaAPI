package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

type ServerConfig struct {
	Port    string `envconfig:"PORT"`
	AppName string `envconfig:"APP_NAME"`
	Url     string `envconfig:"URL"`
}

func CheckEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	var db DatabaseConfig
	if err := envconfig.Process("", &db); err != nil {
		return nil, err
	}
	return &db, nil
}

func GetServerConfig() (*ServerConfig, error) {
	var serv ServerConfig
	if err := envconfig.Process("", &serv); err != nil {
		return nil, err
	}
	return &serv, nil
}
