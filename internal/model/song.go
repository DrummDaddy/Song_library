package model

import "time"

type Song struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Artist    string    `gorm:"type:varchar(255);not null"`
	Album     string    `gorm:"type:varchar(255);not null"`
	Year      int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAT time.Time `gorm:"autoUpdateTime"`
}
