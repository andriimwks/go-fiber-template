package repository

import (
	"github.com/andriimwks/go-fiber-template/internal/models"
	"gorm.io/gorm"
)

type Users interface {
	Create(user models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	ExistsWithNickname(nickname string) (bool, error)
	ExistsWithEmail(email string) (bool, error)
}

type Repositories struct {
	Users Users
}

func New(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: newUserRepository(db),
	}
}
