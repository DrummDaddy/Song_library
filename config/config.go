package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort     string `mapstructure:"SERVER_PORT"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	ExternalAPIURL string `mapstructure:"EXTERNAL_API_URL"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
}

// LoadConfig считывает настройки из файла `.env`.
func LoadConfig() (*Config, error) {
	// Загружаем переменные из .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден")
	}

	// Конфигурация через Viper
	viper.AutomaticEnv() // считываем из окружения
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("EXTERNAL_API_URL", "http://localhost:3000/info")
	viper.SetDefault("LOG_LEVEL", "info")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
