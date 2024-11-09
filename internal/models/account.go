package models

import (
	"gorm.io/gorm"
)

type AccountDB struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Token     string `gorm:"unique"`
	IsDefault bool   `gorm:"default:false"`
}

type Account struct {
	Data struct {
		Username           string `json:"username"`
		Email              string `json:"email"`
		Subscribed         bool   `json:"subscribed"`
		Status             string `json:"status"`
		Badges             []any  `json:"badges"`
		Gems               int    `json:"gems"`
		AchievementsPoints int    `json:"achievements_points"`
		Banned             bool   `json:"banned"`
		BanReason          string `json:"ban_reason"`
	} `json:"data"`
}
