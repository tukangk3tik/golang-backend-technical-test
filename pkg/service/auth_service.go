package service

import (
	"crypto/md5"
	"fmt"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/repository"
)

type AuthService interface {
	Login(email string, pass string) interface{}
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(email string, pass string) interface{} {
	passMd5 := fmt.Sprintf("%x", md5.Sum([]byte(pass)))
	res, err := s.userRepo.Login(email, passMd5)

	if err != nil {
		return false
	}
	return res
}
