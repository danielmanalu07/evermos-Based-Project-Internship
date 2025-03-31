package models

import "time"

type LogProduk struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IdProduk      int       `json:"id_produk" gorm:"type:int;not null"`
	NamaProduk    string    `json:"nama_produk" gorm:"type:varchar(255);not null"`
	Slug          string    `json:"slug" gorm:"type:varchar(255);not null"`
	HargaReseller string    `json:"harga_reseller" gorm:"type:varchar(255);not null"`
	HargaKonsumen string    `json:"harga_konsumen" gorm:"type:varchar(255);not null"`
	Deskripsi     string    `json:"deskripsi" gorm:"type:text;not null"`
	IdToko        uint      `json:"id_toko" gorm:"index;not null"`
	Toko          Toko      `json:"-" gorm:"foreignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdCategory    uint      `json:"id_category" gorm:"index;not null"`
	Category      Category  `json:"-" gorm:"foreignKey:IdCategory;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
