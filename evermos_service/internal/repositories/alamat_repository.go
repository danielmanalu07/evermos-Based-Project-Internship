package repositories

import "evermos-app/internal/models"

type AlamatRepository interface {
	Create(alamat *models.Alamat) error
	FindByUserId(userID uint) ([]models.Alamat, error)
	FindById(id uint) (*models.Alamat, error)
	Update(alamat *models.Alamat) error
	Delete(id uint) error
}
