package config

import "gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"

// Set config for global value in port
var (
	Port   = helper.SetEnv("APP_PORT", "9090")
	DbHost = helper.SetEnv("DB_HOST", "localhost")
	DbPort = helper.SetEnv("DB_PORT", "3306")
	DbName = helper.SetEnv("DB_NAME", "golang_test_db")
	DbUser = helper.SetEnv("DB_USER", "root")
	DbPass = helper.SetEnv("DB_PASS", "pass123")
	JwtKey = helper.SetEnv("JWT_KEY", "1234567890")
)
