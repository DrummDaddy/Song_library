package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstrucure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_Name"`
}

// LoadConfig считывает настройки из файла `.env`.
func LoadConfig() (*Config, error) {
	// Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env файл не найден")
	}

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")

	// Конфигурация через Viper
	viper.AutomaticEnv() // считываем из окружения

	return &Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}, nil

}
