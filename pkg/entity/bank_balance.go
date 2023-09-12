package entity

import "gorm.io/gorm"

type BankBalance struct {
	gorm.Model
	Balance        int    `json:"balance"`
	BalanceAchieve int    `json:"balance_achieve"`
	Code           string `json:"code"`
	Enable         bool   `json:"enable"`
}
