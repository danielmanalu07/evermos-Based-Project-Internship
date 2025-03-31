package repositoryimpl

import (
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

// Create implements repositories.CategoryRepository.
func (c *categoryRepositoryImpl) Create(category *models.Category) error {
	return c.db.Create(category).Error
}

// FindById implements repositories.CategoryRepository.
func (c *categoryRepositoryImpl) FindById(id uint) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// FindAll implements repositories.CategoryRepository.
func (c *categoryRepositoryImpl) FindAll() ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Update implements repositories.CategoryRepository.
func (c *categoryRepositoryImpl) Update(category *models.Category) error {
	return c.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(category).Error
}

// Delete implements repositories.CategoryRepository.
func (c *categoryRepositoryImpl) Delete(id uint) error {
	return c.db.Delete(&models.Category{}, id).Error
}
