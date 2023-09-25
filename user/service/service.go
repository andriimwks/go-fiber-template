package service

import (
	"time"

	"github.com/andriimwks/go-fiber-template/config"
	"github.com/andriimwks/go-fiber-template/models"
	"github.com/andriimwks/go-fiber-template/user"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	config   *config.Config
	repo     user.Repository
	validate *validator.Validate
}

func New(cfg *config.Config, repo user.Repository) user.Service {
	return &service{cfg, repo, validator.New()}
}

func (s *service) SignIn(form user.SignInForm) (*user.TokenPair, error) {
	if err := s.validate.Struct(&form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "email":
				return nil, user.ErrInvalidEmail
			}
		}
	}

	u, err := s.repo.GetByEmail(form.Email)
	if err != nil {
		return nil, user.ErrIncorrectEmailOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)); err != nil {
		return nil, user.ErrIncorrectEmailOrPassword
	}

	return s.createTokenPair(*u)
}

func (s *service) SignUp(form user.SignUpForm) (*user.TokenPair, error) {
	if err := s.validate.Struct(&form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "first_name":
				return nil, user.ErrEmptyFirstName
			case "nickname":
				return nil, user.ErrInvalidNickname
			case "email":
				return nil, user.ErrInvalidEmail
			case "password":
				return nil, user.ErrInvalidPassword
			}
		}
	}

	if exists, _ := s.repo.ExistsWithEmail(form.Email); exists {
		return nil, user.ErrEmailIsTaken
	}

	if exists, _ := s.repo.ExistsWithNickname(form.Nickname); exists {
		return nil, user.ErrNicknameIsTaken
	}

	pwdHash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)

	u, err := s.repo.Create(models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Nickname:  form.Nickname,
		Email:     form.Email,
		Password:  string(pwdHash),
	})
	if err != nil {
		return nil, user.ErrUserExists
	}

	return s.createTokenPair(*u)
}

func (s *service) ValidateTokens(tokens user.TokenPair) (*user.TokenPair, error) {
	atclaims, err := s.parseToken(tokens.AccessToken)
	if err != nil {
		rtclaims, err := s.parseToken(tokens.RefreshToken)
		if err != nil {
			return nil, err
		}

		id, _ := rtclaims["sub"].(float64)
		sub, err := s.repo.GetByID(uint(id))
		if err != nil {
			return nil, err
		}

		return s.createTokenPair(*sub)
	}

	id, _ := atclaims["sub"].(float64)
	sub, err := s.repo.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &user.TokenPair{Subject: *sub, MustUpdate: false}, nil
}

func (s *service) parseToken(raw string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) { return []byte(s.config.Env.JwtSigningKey), nil })
	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	}
	return nil, err
}

func (s *service) createAccessToken(sub models.User) (string, error) {
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

func (s *service) createRefreshToken(id uint) (string, error) {
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

func (s *service) createTokenPair(sub models.User) (*user.TokenPair, error) {
	accToken, err := s.createAccessToken(sub)
	if err != nil {
		return nil, err
	}

	refToken, err := s.createRefreshToken(sub.ID)
	if err != nil {
		return nil, err
	}

	return &user.TokenPair{
		AccessToken:  accToken,
		RefreshToken: refToken,
		Subject:      sub,
		MustUpdate:   true,
	}, nil
}
