package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/controller"
)

func InitRoutes(userController *controller.UserController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index).Methods(http.MethodGet)
	router.HandleFunc("/user", userController.Save).Methods(http.MethodPost)
	router.HandleFunc("/user", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/user/{id:[0-9]+}", userController.FindById).Methods(http.MethodGet)
	return router
}
