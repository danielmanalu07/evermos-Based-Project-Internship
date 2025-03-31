package models

import "time"

type Transaksi struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IdUser           uint      `json:"id_user" gorm:"index;not null"`
	User             User      `json:"-" gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AlamatPengiriman uint      `json:"alamat_pengiriman" gorm:"index;not null"`
	Alamat           Alamat    `json:"-" gorm:"foreignKey:AlamatPengiriman;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HargaTotal       int64     `json:"harga_total" gorm:"type:bigint;not null"`
	KodeInvoice      string    `json:"kode_invoice" gorm:"type:varchar(255);not null"`
	MethodBayar      string    `json:"method_bayar" gorm:"type:varchar(255);not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type DetailTransaksi struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IdTrx       uint      `json:"id_trx" gorm:"index;not null"`
	Transaksi   Transaksi `json:"-" gorm:"foreignKey:IdTrx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdLogProduk uint      `json:"id_log_produk" gorm:"index;not null"`
	LogProduk   LogProduk `json:"-" gorm:"foreignKey:IdLogProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdToko      uint      `json:"id_toko" gorm:"index;not null"`
	Toko        Toko      `json:"-" gorm:"foreignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Kuantitas   int       `json:"kuantitas" gorm:"type:int;not null"`
	HargaTotal  int64     `json:"harga_total" gorm:"type:bigint;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TransactionFilter struct {
	KodeInvoice string `json:"kode_invoice" query:"kode_invoice"`
	MethodBayar string `json:"method_bayar" query:"method_bayar"`
}
