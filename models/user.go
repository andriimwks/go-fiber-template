package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Nickname  string `json:"nickname" gorm:"not null; unique"`
	Email     string `json:"email" gorm:"not null; unique"`
	Password  string `json:"password" gorm:"not null"`
	IsBanned  bool   `json:"is_banned" gorm:"not null; default:0"`
	BanReason string `json:"ban_reason"`
}

func (u *User) Map() fiber.Map {
	return fiber.Map{
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"nickname":   u.Nickname,
		"email":      u.Nickname,
		"is_banned":  u.IsBanned,
		"ban_reason": u.BanReason,
	}
}
