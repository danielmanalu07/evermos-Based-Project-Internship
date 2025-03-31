package repositories

import "evermos-app/internal/models"

type TransactionRepository interface {
	CreateTransaksi(transaksi *models.Transaksi) error
	CreateLogProduk(log *models.LogProduk) (*models.LogProduk, error)
	CreateDetailTransaksi(detail *models.DetailTransaksi) error
	FindTransaksiById(id uint) (*models.Transaksi, error)
	FindDetailsByTransaksiId(transaksiID uint) ([]models.DetailTransaksi, error)
	FindTransaksisByUserId(userID uint) ([]models.Transaksi, error)
	FindTransaksisWithPaginationAndFilter(userID uint, page, limit int, filter *models.TransactionFilter) ([]models.Transaksi, int64, error)
}
