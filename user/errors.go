package user

import "errors"

var (
	ErrEmptyFirstName           = errors.New("empty first name")
	ErrInvalidEmail             = errors.New("invalid email")
	ErrInvalidPassword          = errors.New("invalid password")
	ErrInvalidNickname          = errors.New("invalid nickname")
	ErrEmailIsTaken             = errors.New("email is already taken")
	ErrNicknameIsTaken          = errors.New("nickname is already taken")
	ErrUserExists               = errors.New("such user already exists")
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
)
