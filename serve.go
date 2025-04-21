package main

import (
	"fmt"
	"gitlab.com/tukangk3tik_/golang-backend-technical-test/config"
	"gitlab.com/tukangk3tik_/golang-backend-technical-test/pkg/routes"
)

func main() {
	defer config.CloseDatabaseConnection()

	r := routes.ProvideRoutes()
	port := fmt.Sprintf(":%s", config.Port)
	err := r.Run(port)
	if err != nil {
		panic("error http")
	}
}

func init() {
	config.InitDatabase()
	config.PrePopulateDatabase()
}
