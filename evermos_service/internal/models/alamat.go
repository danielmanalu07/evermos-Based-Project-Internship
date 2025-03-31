package models

import "time"

type Alamat struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IdUser       uint      `json:"id_user" gorm:"index;not null"`
	User         User      `json:"-" gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	JudulAlamat  string    `json:"judul_alamat" gorm:"type:varchar(255);not null"`
	NamaPenerima string    `json:"nama_penerima" gorm:"type:varchar(255);not null"`
	NoTelp       string    `json:"no_telp" gorm:"type:varchar(255);unique;not null"`
	DetailAlamat string    `json:"detail_alamat" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
