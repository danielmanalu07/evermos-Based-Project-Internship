package repositoryimpl

import (
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) Create(user *models.User) error {
	return u.db.Create(user).Error
}

// EmailExists implements repositories.UserRepository.
func (u *userRepositoryImpl) EmailExists(email string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// PhoneExists implements repositories.UserRepository.
func (u *userRepositoryImpl) PhoneExists(phone string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("no_telp = ?", phone).Count(&count)
	return count > 0
}

// FindByEmail implements repositories.UserRepository.
func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindById implements repositories.UserRepository.
func (u *userRepositoryImpl) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Updaet implements repositories.UserRepository.
func (u *userRepositoryImpl) Update(user *models.User) error {
	return u.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// RevokeToken implements repositories.UserRepository.
func (u *userRepositoryImpl) RevokeToken(token string) error {
	return u.db.Create(&models.RevokedToken{Token: token}).Error
}

// IsTokenRevoked implements repositories.UserRepository.
func (u *userRepositoryImpl) IsTokenRevoked(token string) bool {
	var count int64
	u.db.Model(&models.RevokedToken{}).Where("token = ?", token).Count(&count)
	return count > 0
}

// Delete implements repositories.UserRepository.
func (u *userRepositoryImpl) Delete(id uint) error {
	return u.db.Delete(&models.User{}, id).Error
}
