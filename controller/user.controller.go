package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/helper"
	"rest-api/model"
	"rest-api/service"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func New(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.JsonResponseError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := u.userService.FindById(id)
	if err != nil {
		helper.JsonResponseError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	helper.JsonResponseSuccess(writer, user, http.StatusOK)
}

func (u *UserController) FindAll(writer http.ResponseWriter, request *http.Request) {
	users, err := u.userService.FindAll()
	if err != nil {
		helper.JsonResponseError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	helper.JsonResponseSuccess(writer, users, http.StatusOK)
}

func (u *UserController) Save(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("Error json %v", err.Error())
		helper.JsonResponseError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	save, errSave := u.userService.Save(user)
	if errSave != nil {
		log.Printf("Error save %v", errSave.Error())
		helper.JsonResponseError(writer, errSave.Error(), http.StatusBadRequest)
		return
	}
	helper.JsonResponseSuccess(writer, save, http.StatusCreated)
}
