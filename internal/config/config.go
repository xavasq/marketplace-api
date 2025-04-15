package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	JWTSecret  string
}

func LoadEnv() (*Config, error) {
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка при загрузке файла .env: %w", err)
	}

	config := &Config{
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBPort:     viper.GetString("DB_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
	}

	if config.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD не указан в конфигурации")
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET не указан в конфигурации")
	}
	return config, nil
}
