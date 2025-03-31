package models

import "time"

type RevokedToken struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Token     string    `json:"token" gorm:"type:varchar(255);not null;unique"`
	RevokedAt time.Time `json:"revoked_at" gorm:"autoCreateTime"`
}
