package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/inject"
)

func (r Routes) setUserRoutes(rg *gin.RouterGroup) {
	authRoutes := rg.Group("/user")

	authRoutes.Use(inject.AuthMiddleware)
	{
		authRoutes.GET("/balance", inject.UserBalanceController.GetBalance)
		authRoutes.POST("/balance/top-up", inject.UserBalanceController.TopUpBalance)
		authRoutes.POST("/balance/transfer", inject.UserBalanceController.Transfer)
	}
}
