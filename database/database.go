package database

import (
	"Song_library/config"
	"Song_library/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB подключается к базе данных и применяет миграции
func InitDB(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	//Автоматическое применение миграций
	if err := DB.AutoMigrate(&model.Song{}); err != nil {
		log.Fatalf("Не удалось применить миграции: %v", err)
	}

	log.Println("Успешное подключение к базе данных")
}
