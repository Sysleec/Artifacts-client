package dbmodels

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Token     string `gorm:"unique"`
	IsDefault bool   `gorm:"default:false"`
}
