package main

import (
	"log"
	"net/http"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/repository"
	"rest-api/routes"
	"rest-api/service"
)

// @title 	User Service API
// @version	1.0
// @description A User service API in Go using Mux framework

// @contact.name Leo Liscano
// @contact.email andresliscanoa@gmail.com

// @host 	localhost:3000
// @schemes http
func main() {
	db := config.InitDatabase()
	userRepository := repository.New(db)
	userService := service.New(userRepository)
	userController := controller.New(userService)
	router := routes.InitRoutes(userController)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}
