package dbmodels

import "gorm.io/gorm"

type Maps struct {
	gorm.Model
	Name string
	Skin string
	X    int `gorm:"uniqueIndex:idx_x_y"`
	Y    int `gorm:"uniqueIndex:idx_x_y"`
	Type string
	Code string
}
