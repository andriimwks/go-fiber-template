package service

import (
	"github.com/andriimwks/go-fiber-template/internal/config"
	"github.com/andriimwks/go-fiber-template/internal/models"
	"github.com/andriimwks/go-fiber-template/internal/repository"
)

type SignInInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUpInput struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type TokenPair struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	Subject      models.User `json:"subject"`
	MustUpdate   bool        `json:"must_update"`
}

type Users interface {
	SignIn(in SignInInput) (*TokenPair, error)
	SignUp(in SignUpInput) (*TokenPair, error)
	ValidateTokens(tokens TokenPair) (*TokenPair, error)
}

type Services struct {
	Users Users
}

func New(cfg *config.Config, repos *repository.Repositories) *Services {
	return &Services{
		Users: newUserService(cfg, repos),
	}
}
