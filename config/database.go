package config

import (
	"fmt"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbConn = InitDatabase()

func InitDatabase() *gorm.DB {
	dbHost := DbHost
	dbName := DbName
	dbUser := DbUser
	dbPass := DbPass
	dbPort := DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(entity.User{}, entity.UserBalance{}, entity.UserBalanceHistory{}, entity.BankBalance{}, entity.BankBalanceHistory{})
	if err != nil {
		panic("Fail when migrating")
	}

	return db
}

func PrePopulateDatabase() {
	var count int64
	DbConn.Model(&entity.User{}).Unscoped().Count(&count)

	if count == 0 {
		user := []entity.User{
			{
				Username: "felix123",
				Email:    "felix123@mail.com",
				Password: "e10adc3949ba59abbe56e057f20f883e", //123456
			},
			{
				Username: "yan321",
				Email:    "yan321@mail.com",
				Password: "c33367701511b4f6020ec61ded352059", //654321
			},
		}
		err := DbConn.Create(&user).Error
		if err != nil {
			panic("Fail to prepopulate User")
		}

		userBalance := make([]entity.UserBalance, 0)
		for _, u := range user {
			balance := entity.UserBalance{
				UserId:         u.ID,
				Balance:        100,
				BalanceAchieve: 0,
			}
			userBalance = append(userBalance, balance)
		}

		err = DbConn.Create(&userBalance).Error
		if err != nil {
			panic("Fail to prepopulate User Balance")
		}
	}
}

func CloseDatabaseConnection() {
	dbSql, err := DbConn.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	err = dbSql.Close()
	if err != nil {
		panic(err)
	}
}
