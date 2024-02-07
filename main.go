package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/database"
	"rest-api/routes"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	database.InitDatabase()
	routes.InitRoutes(router)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}
