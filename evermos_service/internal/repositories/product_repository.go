package repositories

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	CreateFotoProduk(foto *models.FotoProduk) error
	FindByIdWithFotos(productID uint) (*models.Product, error)
	FindByTokoId(tokoID uint) ([]models.Product, error)
	FindAll() ([]models.Product, error)
	FindById(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
	FindAllWithPaginationAndFilter(userID uint, page, limit int, filter *dtos.ProductFilter) ([]models.Product, int64, error)
}
