package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"
)

type Routes struct {
	router *gin.Engine
}

func ProvideRoutes() *gin.Engine {
	apiPath := "/api"

	r := Routes{router: gin.Default()}

	r.router.GET("/api", func(c *gin.Context) {
		c.JSON(200, "API v1")
	})
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.BuildFailResponse("not found", helper.EmptyObject{}))
	})

	path := r.router.Group(apiPath)
	r.setAuthRoutes(path)
	r.setUserRoutes(path)

	return r.router
}

func (r Routes) Run(addr string) error {
	return r.router.Run(addr)
}
