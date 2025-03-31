package repositoryimpl

import (
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repositories.TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

// CreateTransaksi implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) CreateTransaksi(transaksi *models.Transaksi) error {
	return t.db.Create(transaksi).Error
}

// CreateLogProduk implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) CreateLogProduk(log *models.LogProduk) (*models.LogProduk, error) {
	if err := t.db.Create(log).Error; err != nil {
		return nil, err
	}
	return log, nil
}

// CreateDetailTransaksi implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) CreateDetailTransaksi(detail *models.DetailTransaksi) error {
	return t.db.Create(detail).Error
}

// FindTransaksiById implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) FindTransaksiById(id uint) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	if err := t.db.Preload("Alamat").First(&transaksi, id).Error; err != nil {
		return nil, err
	}
	return &transaksi, nil
}

// FindDetailsByTransaksiId implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) FindDetailsByTransaksiId(transaksiID uint) ([]models.DetailTransaksi, error) {
	var details []models.DetailTransaksi
	if err := t.db.Preload("LogProduk").Preload("Toko").Where("id_trx = ?", transaksiID).Find(&details).Error; err != nil {
		return nil, err
	}
	return details, nil
}

// FindTransaksisByUserId implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) FindTransaksisByUserId(userID uint) ([]models.Transaksi, error) {
	var transaksis []models.Transaksi
	if err := t.db.Preload("Alamat").Where("id_user = ?", userID).Find(&transaksis).Error; err != nil {
		return nil, err
	}
	return transaksis, nil
}

// FindTransaksisWithPaginationAndFilter implements repositories.TransactionRepository.
func (t *transactionRepositoryImpl) FindTransaksisWithPaginationAndFilter(userID uint, page int, limit int, filter *models.TransactionFilter) ([]models.Transaksi, int64, error) {
	var transaksis []models.Transaksi
	var total int64

	// Hitung offset
	offset := (page - 1) * limit

	query := t.db.Where("id_user = ?", userID).Preload("Alamat")

	if filter != nil {
		if filter.KodeInvoice != "" {
			query = query.Where("kode_invoice LIKE ?", "%"+filter.KodeInvoice+"%")
		}
		if filter.MethodBayar != "" {
			query = query.Where("method_bayar = ?", filter.MethodBayar)
		}
	}

	if err := query.Model(&models.Transaksi{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&transaksis).Error; err != nil {
		return nil, 0, err
	}

	return transaksis, total, nil
}
