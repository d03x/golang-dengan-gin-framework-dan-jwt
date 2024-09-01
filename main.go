package main

import (
	"dadandev.com/golang-dasar/internal/repositories"
	"dadandev.com/golang-dasar/internal/routes"
	"dadandev.com/golang-dasar/internal/service"
	"github.com/gin-gonic/gin"
)

// run main entry point
func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//inject semua dependensi kesini
	userRepo := repositories.NewUser()
	userService := service.NewAuth(userRepo)
	//grouping router ke beberapa handelr
	routes.NewUser(router, userService)
	router.Run()
}
