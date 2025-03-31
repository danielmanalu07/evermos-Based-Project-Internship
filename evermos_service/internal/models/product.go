package models

import "time"

type Product struct {
	ID            uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaProduk    string       `json:"nama_produk" gorm:"type:varchar(255);not null"`
	Slug          string       `json:"slug" gorm:"type:varchar(255);not null"`
	HargaReseller string       `json:"harga_reseller" gorm:"type:varchar(255);not null"`
	HargaKonsumen string       `json:"harga_konsumen" gorm:"type:varchar(255);not null"`
	Stok          int          `json:"stok" gorm:"type:int;not null"`
	Deskripsi     string       `json:"deskripsi" gorm:"type:text;not null"`
	IdToko        uint         `json:"id_toko" gorm:"index;not null"`
	Toko          Toko         `json:"-" gorm:"foreignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FotoProduk    []FotoProduk `json:"foto_produk" gorm:"foreignKey:IdProduk"`
	IdCategory    uint         `json:"id_category" gorm:"index;not null"`
	Category      Category     `json:"-" gorm:"foreignKey:IdCategory;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProductFilter struct {
	NamaProduk string `json:"nama_produk" query:"nama_produk"`
	IdToko     *uint  `json:"id_toko" query:"id_toko"`
	IdCategory *uint  `json:"id_category" query:"id_category"`
}
