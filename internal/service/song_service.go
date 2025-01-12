package service

import (
	"Song_library/internal/model"
	"database/sql"
)

type SongService struct {
	DB *sql.DB
}

func (s *SongService) GetAllSongs() ([]model.Song, error) {
	rows, err := s.DB.Query("SELECT * FROM somgs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		if err := rows.Scan(
			&song.ID, &song.GroupName, &song.SongName,
			&song.ReleaseDate, &song.Lyrics, &song.YoutubeLink,
			&song.CreatedAt, &song.UpdatedAt,
		); err != nil {
			return nil, err
		}
		songs = append(songs, song)

	}
	return songs, nil

}

func (s *SongService) CreatedSong(song model.Song) error {
	_, err := s.DB.Exec(
		`INSERT INTO songs(group_name, song_name, release_date, lyrics, youtube_link)
		VALUES ($1, $2, $3, $4, $5)`,
		song.GroupName, song.SongName, song.ReleaseDate, song.Lyrics, song.YoutubeLink)
	return err
}
