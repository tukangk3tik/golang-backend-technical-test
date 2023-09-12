package inject

import (
	"gitlab.com/tukangk3tik_/privyid-golang-test/config"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/controller"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/middleware"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/repository"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/service"
)

var (
	// repo
	UserRepo           = repository.NewUserRepository(config.DbConn)
	UserBalanceRepo    = repository.NewUserBalanceRepository(config.DbConn)
	JwtInstanceService = service.NewJwtService(UserRepo)

	// service
	AuthService        = service.NewAuthService(UserRepo)
	UserBalanceService = service.NewUserBalanceService(UserBalanceRepo, UserRepo)

	//controller
	AuthController        = controller.NewAuthController(AuthService, JwtInstanceService)
	UserBalanceController = controller.NewUserBalanceController(UserBalanceService)

	AuthMiddleware = middleware.AuthJwt(JwtInstanceService)
)
