package main

import (
	"Song_library/config"
	"Song_library/database"
	"Song_library/internal/handler"
	"Song_library/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Онлайн Библиотека Песен API
// @version 1.0
// @description REST API для управления песнями
// @termsOfService http://example.com/terms/

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	database.InitDB(cfg)

	r := gin.Default()

	//Инициализация сервисов и обработчиков
	songService := &service.SongService{}
	songHandler := &handler.SongHandler{Service: songService}

	// Маршруты
	r.POST("/songs", songHandler.CreateSong)
	r.GET("/songs", songHandler.GetSongs)
	r.GET("/songs/:id", songHandler.GetSongByID)
	r.PUT("/songs", songHandler.UpdateSong)
	r.DELETE("/songs/:id", songHandler.DeleteSong)

	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
