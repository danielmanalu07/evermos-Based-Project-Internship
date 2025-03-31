package models

import "time"

type Category struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaCategory string    `json:"nama_category" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
