package migrations

import (
	"evermos-app/internal/models"

	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Alamat{},
		&models.Toko{},
		&models.RevokedToken{},
		&models.Category{},
		&models.Product{},
		&models.FotoProduk{},
		&models.LogProduk{},
		&models.Transaksi{},
		&models.DetailTransaksi{},
	)
}
