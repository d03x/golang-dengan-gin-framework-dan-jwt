package main

import (
	"embed"
	"fmt"

	"dadandev.com/golang-dasar/internal/repositories"
	"dadandev.com/golang-dasar/internal/routes"
	"dadandev.com/golang-dasar/internal/service"
	"github.com/gin-gonic/gin"
)

//go:embed html/*
var content embed.FS

func main() {
	fmt.Print(content)
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
