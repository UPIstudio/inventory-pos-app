package repositories

import (
	"github.com/luthfi/inventory-pos-app/backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Save(user *models.User) error
	UpdateActiveToken(id uint, token string) error
	FindByID(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) Save(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateActiveToken(id uint, token string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("active_token", token).Error
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}
