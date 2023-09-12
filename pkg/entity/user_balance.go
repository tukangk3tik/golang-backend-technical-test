package entity

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserId         uint `json:"user_id"`
	Balance        uint `json:"balance"`
	BalanceAchieve uint `json:"balance_achieve"`
	User           User `json:"user" gorm:"foreignKey:UserId"`
}
