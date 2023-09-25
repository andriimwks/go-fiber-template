package gorm

import (
	"github.com/andriimwks/go-fiber-template/models"
	"github.com/andriimwks/go-fiber-template/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &repository{db}
}

func (r *repository) Create(user models.User) (*models.User, error) {
	return &user, r.db.Create(&user).Error
}

func (r *repository) GetByID(id uint) (*models.User, error) {
	user := new(models.User)
	tx := r.db.Where("id = ?", id).First(user)
	return user, tx.Error
}

func (r *repository) GetByNickname(nickname string) (*models.User, error) {
	user := new(models.User)
	tx := r.db.Where("nickname = ?", nickname).First(user)
	return user, tx.Error
}

func (r *repository) GetByEmail(email string) (*models.User, error) {
	user := new(models.User)
	tx := r.db.Where("email = ?", email).First(user)
	return user, tx.Error
}

func (r *repository) ExistsWithNickname(nickname string) (bool, error) {
	var count int64
	tx := r.db.Where("nickname = ?", nickname).Limit(1).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

func (r *repository) ExistsWithEmail(email string) (bool, error) {
	var count int64
	tx := r.db.Where("email = ?", email).Limit(1).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}
