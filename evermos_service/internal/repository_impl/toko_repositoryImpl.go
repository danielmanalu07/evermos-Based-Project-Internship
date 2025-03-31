package repositoryimpl

import (
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type tokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) repositories.TokoRepository {
	return &tokoRepositoryImpl{db: db}
}

func (t *tokoRepositoryImpl) Create(toko *models.Toko) error {
	return t.db.Create(toko).Error
}

// FindByUserId implements repositories.TokoRepository.
func (t *tokoRepositoryImpl) FindByUserId(userID uint) (*models.Toko, error) {
	var tokos models.Toko
	if err := t.db.Where("id_user = ?", userID).Find(&tokos).Error; err != nil {
		return nil, err
	}
	return &tokos, nil
}

// FindById implements repositories.TokoRepository.
func (t *tokoRepositoryImpl) FindById(id uint) (*models.Toko, error) {
	var toko models.Toko
	if err := t.db.First(&toko, id).Error; err != nil {
		return nil, err
	}
	return &toko, nil
}

// Update implements repositories.TokoRepository.
func (t *tokoRepositoryImpl) Update(toko *models.Toko) error {
	return t.db.Model(&models.Toko{}).Where("id = ?", toko.ID).Updates(toko).Error
}

// Delete implements repositories.TokoRepository.
func (t *tokoRepositoryImpl) Delete(id uint) error {
	return t.db.Delete(&models.Toko{}, id).Error
}
