package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/inject"
)

func (r Routes) setAuthRoutes(rg *gin.RouterGroup) {
	authRoutes := rg.Group("/auth")

	authRoutes.POST("/login", inject.AuthController.Login)

	authRoutes.Use(inject.AuthMiddleware)
	{
		authRoutes.POST("/logout", inject.AuthController.Logout)
	}
}
