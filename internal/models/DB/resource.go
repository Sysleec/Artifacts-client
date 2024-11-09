package dbmodels

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name  string
	Code  string `gorm:"unique"`
	Skill string
	Level int
	Drops []Drop `gorm:"foreignKey:ResourceID"`
}

type Drop struct {
	gorm.Model
	ResourceID  uint
	Code        string
	Rate        int
	MinQuantity int
	MaxQuantity int
}
