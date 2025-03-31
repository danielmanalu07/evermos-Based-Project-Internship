package repositoryimpl

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &productRepositoryImpl{db: db}
}

// Create implements repositories.ProductRepository.
func (p *productRepositoryImpl) Create(product *models.Product) error {
	return p.db.Create(product).Error
}

// CreateFotoProduk implements repositories.ProductRepository.
func (p *productRepositoryImpl) CreateFotoProduk(foto *models.FotoProduk) error {
	return p.db.Create(foto).Error
}

// FindByIdWithFotos implements repositories.ProductRepository.
func (p *productRepositoryImpl) FindByIdWithFotos(productID uint) (*models.Product, error) {
	var product models.Product
	err := p.db.Preload("FotoProduk").First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindAll implements repositories.ProductRepository.
func (p *productRepositoryImpl) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Preload("FotoProduk").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindByTokoId implements repositories.ProductRepository.
func (p *productRepositoryImpl) FindByTokoId(tokoID uint) ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Preload("FotoProduk").Where("id_toko = ?", tokoID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindById implements repositories.ProductRepository.
func (p *productRepositoryImpl) FindById(id uint) (*models.Product, error) {
	var product models.Product
	if err := p.db.Preload("FotoProduk").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Update implements repositories.ProductRepository.
func (p *productRepositoryImpl) Update(product *models.Product) error {
	return p.db.Model(&models.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

// Delete implements repositories.ProductRepository.
func (p *productRepositoryImpl) Delete(id uint) error {
	return p.db.Delete(&models.Product{}, id).Error
}

// FindAllWithPaginationAndFilter implements repositories.ProductRepository.
func (p *productRepositoryImpl) FindAllWithPaginationAndFilter(userID uint, page, limit int, filter *dtos.ProductFilter) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	query := p.db.Joins("JOIN tokos ON tokos.id = products.id_toko").
		Where("tokos.id_user = ?", userID).
		Preload("FotoProduk")

	if filter != nil {
		if filter.NamaProduk != "" {
			query = query.Where("products.nama_produk LIKE ?", "%"+filter.NamaProduk+"%")
		}
		if filter.IdToko != nil {
			query = query.Where("products.id_toko = ?", *filter.IdToko)
		}
		if filter.IdCategory != nil {
			query = query.Where("products.id_category = ?", *filter.IdCategory)
		}
	}

	if err := query.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
