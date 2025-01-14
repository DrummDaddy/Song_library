package service

import (
	"Song_library/database"
	"Song_library/internal/model"
)

// SongsService предоставляет методы работы с песнями
type SongService struct{}

// CreateSong добавляет новую песню в базу данных
func (s *SongService) CreateSong(song *model.Song) (*model.Song, error) {
	if err := database.DB.Create(song).Error; err != nil {
		return nil, err
	}

	return song, nil
}

// GetSongs возвращяет все псени
func (s *SongService) GetSongs() ([]model.Song, error) {
	var songs []model.Song
	if err := database.DB.Find(&songs).Error; err != nil {
		return nil, err
	}

	return songs, nil
}

// GetSongByID возвращяет песню по ID
func (s *SongService) GetSongByID(id uint) (*model.Song, error) {
	var song model.Song
	if err := database.DB.First(&song, id).Error; err != nil {
		return nil, err
	}

	return &song, nil
}

// DeleteSong удаляет песню по ID
func (s *SongService) DeleteSong(id uint) error {
	if err := database.DB.Delete(&model.Song{}, id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateSong обновляет данные песни
func (s *SongService) UpdateSong(song *model.Song) (*model.Song, error) {
	if err := database.DB.Save(song).Error; err != nil {
		return nil, err
	}
	return song, nil
}
