package models

import "time"

type FotoProduk struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IdProduk  uint      `json:"id_produk" gorm:"index;not null"`
	Produk    Product   `json:"-" gorm:"foreignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Url       string    `json:"url" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
