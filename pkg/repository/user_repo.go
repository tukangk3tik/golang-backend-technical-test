package repository

import (
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email string, pass string) (entity.User, error)
	GetUser(userId uint) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
}

type userConnection struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{conn: db}
}

func (db *userConnection) Login(email string, pass string) (user entity.User, err error) {
	err = db.conn.Where("email = ? AND password = ?", email, pass).Take(&user).Error
	return
}

func (db *userConnection) GetUser(userId uint) (user entity.User, err error) {
	err = db.conn.First(&user, userId).Error
	return
}

func (db *userConnection) GetUserByUsername(username string) (user entity.User, err error) {
	err = db.conn.Where("username = ?", username).First(&user).Error
	return
}
