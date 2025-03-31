package repositoryimpl

import (
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type alamatRepositoryImpl struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) repositories.AlamatRepository {
	return &alamatRepositoryImpl{db: db}
}

// Create implements repositories.AlamatRepository.
func (a *alamatRepositoryImpl) Create(alamat *models.Alamat) error {
	return a.db.Create(alamat).Error
}

// FindByUserId implements repositories.AlamatRepository.
func (a *alamatRepositoryImpl) FindByUserId(userID uint) ([]models.Alamat, error) {
	var alamats []models.Alamat
	if err := a.db.Where("id_user = ?", userID).Find(&alamats).Error; err != nil {
		return nil, err
	}
	return alamats, nil
}

// FindById implements repositories.AlamatRepository.
func (a *alamatRepositoryImpl) FindById(id uint) (*models.Alamat, error) {
	var alamat models.Alamat
	if err := a.db.First(&alamat, id).Error; err != nil {
		return nil, err
	}
	return &alamat, nil
}

// Update implements repositories.AlamatRepository.
func (a *alamatRepositoryImpl) Update(alamat *models.Alamat) error {
	return a.db.Model(&models.Alamat{}).Where("id = ?", alamat.ID).Updates(alamat).Error
}

// Delete implements repositories.AlamatRepository.
func (a *alamatRepositoryImpl) Delete(id uint) error {
	return a.db.Delete(&models.Alamat{}, id).Error
}
