package repositories

import "evermos-app/internal/models"

type TokoRepository interface {
	Create(toko *models.Toko) error
	FindByUserId(userID uint) (*models.Toko, error)
	FindById(id uint) (*models.Toko, error)
	Update(toko *models.Toko) error
	Delete(id uint) error
}
