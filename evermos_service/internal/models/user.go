package models

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama         string    `json:"nama" gorm:"type:varchar(255);not null"`
	KataSandi    string    `json:"kata_sandi" gorm:"type:varchar(255);not null"`
	NoTelp       string    `json:"no_telp" gorm:"type:varchar(255);unique;not null"`
	TanggalLahir time.Time `json:"tanggal_lahir" gorm:"type:date;not null"`
	JenisKelamin string    `json:"jenis_kelamin" gorm:"type:varchar(10);not null"`
	Tentang      string    `json:"tentang" gorm:"type:text"`
	Pekerjaan    string    `json:"pekerjaan" gorm:"type:varchar(255)"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	IdProvinsi   string    `json:"id_provinsi" gorm:"type:varchar(255)"`
	IdKota       string    `json:"id_kota" gorm:"type:varchar(255)"`
	IsAdmin      bool      `json:"is_admin" gorm:"type:boolean;default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
