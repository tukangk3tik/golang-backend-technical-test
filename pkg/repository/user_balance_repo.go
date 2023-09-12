package repository

import (
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/request"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/entity"
	"gorm.io/gorm"
)

type UserBalanceRepository interface {
	GetUserBalance(userId uint) (entity.UserBalance, error)
	AddBalance(userBalanceId, amount uint) error
	DeductBalance(userBalanceId, amount uint) error
	CreateBalanceHistory(header request.HeaderDto, dataHistory ...map[string]any) (err error)
}

type balanceConnection struct {
	conn *gorm.DB
}

func NewUserBalanceRepository(db *gorm.DB) UserBalanceRepository {
	return &balanceConnection{conn: db}
}

func (db *balanceConnection) GetUserBalance(userId uint) (userBalance entity.UserBalance, err error) {
	err = db.conn.Where("user_id = ?", userId).First(&userBalance).Error
	return
}

func (db *balanceConnection) AddBalance(userBalanceId, amount uint) (err error) {
	var userBalance entity.UserBalance
	err = db.conn.First(&userBalance, userBalanceId).Error
	if err != nil {
		return
	}

	userBalance.Balance += amount
	err = db.conn.Save(&userBalance).Error
	return
}

func (db *balanceConnection) DeductBalance(userBalanceId, amount uint) (err error) {
	var userBalance entity.UserBalance
	err = db.conn.First(&userBalance, userBalanceId).Error
	if err != nil {
		return
	}

	userBalance.Balance -= amount
	err = db.conn.Save(&userBalance).Error
	return
}

func (db *balanceConnection) CreateBalanceHistory(header request.HeaderDto, dataHistories ...map[string]any) (err error) {
	history := make([]entity.UserBalanceHistory, 0)

	for _, v := range dataHistories {
		item := entity.UserBalanceHistory{
			UserBalanceId: v["userBalanceId"].(uint),
			BalanceBefore: v["balanceBefore"].(uint),
			BalanceAfter:  v["balanceAfter"].(uint),
			Activity:      "", // not sure for what
			Type:          v["type"].(entity.HistoryType),
			Ip:            header.Ip,
			Location:      header.Location,
			UserAgent:     header.UserAgent,
			Author:        header.Author,
		}

		history = append(history, item)
	}

	err = db.conn.Save(&history).Error
	return
}
