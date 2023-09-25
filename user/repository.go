package user

import (
	"github.com/andriimwks/go-fiber-template/models"
)

type Repository interface {
	Create(user models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByNickname(nickname string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	ExistsWithNickname(nickname string) (bool, error)
	ExistsWithEmail(email string) (bool, error)
}
