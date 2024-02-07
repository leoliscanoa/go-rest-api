package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/handler"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/", handler.Index).Methods(http.MethodGet)
	router.HandleFunc("/user", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", handler.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id:[0-9]+}", handler.GetUserById).Methods(http.MethodGet)
}
