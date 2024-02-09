package main

import (
	"log"
	"net/http"
	"rest-api/controller"
	"rest-api/database"
	"rest-api/repository"
	"rest-api/routes"
	"rest-api/service"
)

func main() {
	db := database.InitDatabase()
	userRepository := repository.New(db)
	userService := service.New(userRepository)
	userController := controller.New(userService)
	router := routes.InitRoutes(userController)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}
