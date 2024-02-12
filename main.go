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

func main() {
	db := config.InitDatabase()
	userRepository := repository.New(db)
	userService := service.New(userRepository)
	userController := controller.New(userService)
	router := routes.InitRoutes(userController)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}
