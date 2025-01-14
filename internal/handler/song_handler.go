package handler

import (
	"Song_library/internal/model"
	"Song_library/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	Service *service.SongService
}

// CreateSong обработчик для добавления песни
// @Summary Создать песню
// @Description Добавление новой песни в базу данных
// @Tags Песни
// @Accept json
// @Produce json
// @Param song body model.Song true "Информация о песне"
// @Success 201 {object} model.Song "Успешное создание песни"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Failure 500 {object} map[string]string "Внутренняя ошибка"
// @Router /songs [post]
func (h *SongHandler) CreateSong(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSong, err := h.Service.CreateSong(&song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось созадть песню"})
		return
	}
	c.JSON(http.StatusCreated, createdSong)
}

// GetSongs возвращает список всех песен
// @Summary Список всех песен
// @Description Получение списка всех песен
// @Tags Песни
// @Accept json
// @Produce json
// @Success 200 {array} model.Song "Список песен"
// @Failure 500 {object} map[string]string "Внутренняя ошибка"
// @Router /songs [get]
func (h *SongHandler) GetSongs(c *gin.Context) {
	songs, err := h.Service.GetSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить песни"})
		return
	}
	c.JSON(http.StatusOK, songs)
}

// GetSongsByID возвращает песню по ID
// @Summary Получить песню по ID
// @Description Получение информации о песне по ID
// @Tags Песни
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} model.Song "Информация о песне"
// @Failure 400 {object} map[string]string "Некорректный ID"
// @Failure 404 {object} map[string]string "Песня не найдена"
// @Router /songs/{id} [get]
func (h *SongHandler) GetSongByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}
	song, err := h.Service.GetSongByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}
	c.JSON(http.StatusOK, song)
}

// DeleteSong удаляет песню
// @Summary Удалить песню
// @Description Удаление песни по ID
// @Tags Песни
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} map[string]string "Успешное удаление"
// @Failure 400 {object} map[string]string "Некорректный ID"
// @Failure 500 {object} map[string]string "Ошибка удаления"
// @Router /songs/{id} [delete]
func (h *SongHandler) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := h.Service.DeleteSong(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить песню"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Песня удалена"})
}

// UpdateSong обновляет данные песни
// @Summary Обновить песню
// @Description Обновление информации о песне
// @Tags Песни
// @Accept json
// @Produce json
// @Param song body model.Song true "Информация о песне"
// @Success 200 {object} model.Song "Обновленная информация о песне"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Failure 500 {object} map[string]string "Ошибка обновления"
// @Router /songs [put]
func (h *SongHandler) UpdateSong(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSong, err := h.Service.UpdateSong(&song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить песню"})
		return
	}

	c.JSON(http.StatusOK, updatedSong)
}
