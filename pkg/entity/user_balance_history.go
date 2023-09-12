package entity

import "gorm.io/gorm"

type HistoryType string

const (
	DEBIT  HistoryType = "DEBIT"
	KREDIT HistoryType = "KREDIT"
)

type UserBalanceHistory struct {
	gorm.Model
	UserBalanceId uint        `json:"user_balance_id"`
	BalanceBefore uint        `json:"balance_before"`
	BalanceAfter  uint        `json:"balance_after"`
	Activity      string      `json:"activity"`
	Type          HistoryType `json:"type"`
	Ip            string      `json:"ip"`
	Location      string      `json:"location"`
	UserAgent     string      `json:"user_agent"`
	Author        string      `json:"author"`
	UserBalance   UserBalance `json:"user_balance" gorm:"foreignKey:UserBalanceId"`
}
