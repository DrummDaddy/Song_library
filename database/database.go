package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Ошибка соединения с БД: %v", err)
	}

	//Миграции
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS songs(
		id SERIAL PRIMARY KEY,
		group_name VARCHAR(255) NOT NULL,
		song_name VARCHAR(255) NOT NULL,
		release_date DATE, 
		lyrics TEXT,
		youtube_link TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`); err != nil {
		log.Fatalf("Миграция не выполнена: %v", err)
	}
	return db
}
