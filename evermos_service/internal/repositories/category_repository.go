package repositories

import "evermos-app/internal/models"

type CategoryRepository interface {
	Create(category *models.Category) error
	FindById(id uint) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}
