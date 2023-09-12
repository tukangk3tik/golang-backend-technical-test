package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/tukangk3tik_/privyid-golang-test/config"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/response/auth"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/repository"
	"time"
)

var whiteListToken = make([]string, 0)

type JwtService interface {
	GenerateToken(userId uint) auth.TokenDto
	ValidateToken(token string) (any, error)
	RemoveToken(token string) (err error)
}

type Jwt struct {
	userRepo repository.UserRepository
}

func NewJwtService(userRepo repository.UserRepository) JwtService {
	return &Jwt{
		userRepo: userRepo,
	}
}

func (j *Jwt) GenerateToken(userId uint) (token auth.TokenDto) {
	tokenType := "Bearer"
	expiresAt := time.Now().AddDate(0, 0, 30)
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["exp"] = expiresAt.Unix()
	claims["userid"] = userId

	at, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JwtKey))
	if err != nil {
		panic(err)
	}

	token = auth.TokenDto{
		TokenType:   tokenType,
		ExpiresIn:   expiresAt.Unix(),
		AccessToken: at,
	}

	// add token to whitelist
	whiteListToken = append(whiteListToken, at)
	return
}

func (j *Jwt) ValidateToken(token string) (claim any, err error) {
	check := helper.IsExistsInSlice[string](whiteListToken, token)
	if !check {
		return nil, fmt.Errorf("unauthorized token")
	}

	var t *jwt.Token
	t, err = jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(config.JwtKey), nil
	})

	if err != nil {
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		err = fmt.Errorf("token invalid")
		return
	}

	claim = uint(claims["userid"].(float64))
	return
}

func (j *Jwt) RemoveToken(token string) (err error) {
	_, err = j.ValidateToken(token)
	whiteListToken = helper.RemoveFromSlice[string](whiteListToken, token)
	return
}
