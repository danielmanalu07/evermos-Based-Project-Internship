package repositories

import "evermos-app/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	EmailExists(email string) bool
	PhoneExists(phone string) bool
	FindByEmail(email string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	Update(user *models.User) error
	RevokeToken(token string) error
	IsTokenRevoked(token string) bool
	Delete(id uint) error
}
