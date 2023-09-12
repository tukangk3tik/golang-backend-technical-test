package entity

import "gorm.io/gorm"

type BankBalanceHistory struct {
	gorm.Model
	BankBalanceId uint        `json:"bank_balance_id"`
	BalanceBefore int         `json:"balance_before"`
	BalanceAfter  int         `json:"balance_after"`
	Activity      int         `json:"activity"`
	Type          string      `json:"type"`
	Ip            string      `json:"ip"`
	Location      string      `json:"location"`
	UserAgent     string      `json:"user_agent"`
	Author        string      `json:"author"`
	UserBalance   UserBalance `json:"user_balance" gorm:"foreignKey:BankBalanceId"`
}
