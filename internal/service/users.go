package service

import (
	"errors"
	"time"

	"github.com/andriimwks/go-fiber-template/internal/config"
	"github.com/andriimwks/go-fiber-template/internal/models"
	"github.com/andriimwks/go-fiber-template/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	errEmailIsTaken             = errors.New("email is already taken")
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type userService struct {
	config   *config.Config
	repos    *repository.Repositories
	validate *validator.Validate
}

func newUserService(cfg *config.Config, repos *repository.Repositories) Users {
	return &userService{cfg, repos, validator.New()}
}

func (s *userService) SignIn(in SignInInput) (*TokenPair, error) {
	if err := s.validate.Struct(&in); err != nil {
		return nil, err
	}

	u, err := s.repos.Users.GetByEmail(in.Email)
	if err != nil {
		return nil, errIncorrectEmailOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)); err != nil {
		return nil, errIncorrectEmailOrPassword
	}

	return s.createTokenPair(*u)
}

func (s *userService) SignUp(in SignUpInput) (*TokenPair, error) {
	if err := s.validate.Struct(&in); err != nil {
		return nil, err
	}

	if exists, _ := s.repos.Users.ExistsWithEmail(in.Email); exists {
		return nil, errEmailIsTaken
	}

	pwdHash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	u, err := s.repos.Users.Create(models.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  string(pwdHash),
	})
	if err != nil {
		return nil, err
	}

	return s.createTokenPair(*u)
}

func (s *userService) ValidateTokens(tokens TokenPair) (*TokenPair, error) {
	atclaims, err := s.parseToken(tokens.AccessToken)
	if err != nil {
		rtclaims, err := s.parseToken(tokens.RefreshToken)
		if err != nil {
			return nil, err
		}

		id, _ := rtclaims["sub"].(float64)
		sub, err := s.repos.Users.GetByID(uint(id))
		if err != nil {
			return nil, err
		}

		return s.createTokenPair(*sub)
	}

	id, _ := atclaims["sub"].(float64)
	sub, err := s.repos.Users.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &TokenPair{Subject: *sub, MustUpdate: false}, nil
}

func (s *userService) parseToken(raw string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) { return []byte(s.config.Env.JwtSigningKey), nil })
	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	}
	return nil, err
}

func (s *userService) createAccessToken(sub models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = float64(sub.ID)
	claims["exp"] = time.Now().Add(s.config.Auth.AccessTokenLifetime).Unix()

	t, err := token.SignedString([]byte(s.config.Env.JwtSigningKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *userService) createRefreshToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = float64(id)
	claims["exp"] = time.Now().Add(s.config.Auth.RefreshTokenLifetime).Unix()

	t, err := token.SignedString([]byte(s.config.Env.JwtSigningKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *userService) createTokenPair(sub models.User) (*TokenPair, error) {
	accToken, err := s.createAccessToken(sub)
	if err != nil {
		return nil, err
	}

	refToken, err := s.createRefreshToken(sub.ID)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accToken,
		RefreshToken: refToken,
		Subject:      sub,
		MustUpdate:   true,
	}, nil
}
