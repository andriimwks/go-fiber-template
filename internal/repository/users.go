package repository

import (
	"github.com/andriimwks/go-fiber-template/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) Users {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.User) (*models.User, error) {
	return &user, r.db.Create(&user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	user := new(models.User)
	tx := r.db.Where("id = ?", id).First(user)
	return user, tx.Error
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	user := new(models.User)
	tx := r.db.Where("email = ?", email).First(user)
	return user, tx.Error
}

func (r *userRepository) ExistsWithNickname(nickname string) (bool, error) {
	var count int64
	tx := r.db.Where("nickname = ?", nickname).Limit(1).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

func (r *userRepository) ExistsWithEmail(email string) (bool, error) {
	var count int64
	tx := r.db.Where("email = ?", email).Limit(1).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}
