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
}

func LoadEnv() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("ошибка при загрузке файла .env: %v", err)
	}

	return &Config{
		DBName:     viper.GetString("DBName"),
		DBUser:     viper.GetString("DBUser"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBPort:     viper.GetString("DB_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
	}
}
