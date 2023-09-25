package user

import (
	"github.com/andriimwks/go-fiber-template/models"
)

type Service interface {
	SignIn(form SignInForm) (*TokenPair, error)
	SignUp(form SignUpForm) (*TokenPair, error)
	ValidateTokens(tokens TokenPair) (*TokenPair, error)
}

type TokenPair struct {
	AccessToken, RefreshToken string
	Subject                   models.User
	MustUpdate                bool
}
