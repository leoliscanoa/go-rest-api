package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "rest-api/dto"
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

// FindById			godoc
// @Summary			find a user
// @Description		Find a user by it\'s unique identification.
// @Param			id path string true "user unique identification"
// @Produce			application/json
// @Tags			users
// @Success			201  {object}  dto.ResponseDto[model.User]
// @Failure      	400  {object}  dto.ResponseDto[string]
// @Failure      	401  {object}  dto.ResponseDto[string]
// @Failure      	500  {object}  dto.ResponseDto[string]
// @Router			/user/{id} [get]
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

// FindAll			godoc
// @Summary			find all users
// @Description		Find all registered users
// @Produce			application/json
// @Tags			users
// @Success			201  {object}  dto.ResponseDto[[]model.User]
// @Failure      	400  {object}  dto.ResponseDto[string]
// @Failure      	401  {object}  dto.ResponseDto[string]
// @Failure      	500  {object}  dto.ResponseDto[string]
// @Router			/user [get]
func (u *UserController) FindAll(writer http.ResponseWriter, request *http.Request) {
	users, err := u.userService.FindAll()
	if err != nil {
		helper.JsonResponseError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	helper.JsonResponseSuccess(writer, users, http.StatusOK)
}

// Save 			godoc
// @Summary			Create user
// @Description		Save user data in Db.
// @Param			request body model.User true "Create user"
// @Produce			json
// @Accept			json
// @Tags			users
// @Success			201  {object}  dto.ResponseDto[model.User]
// @Failure      	400  {object}  dto.ResponseDto[string]
// @Failure      	401  {object}  dto.ResponseDto[string]
// @Failure      	500  {object}  dto.ResponseDto[string]
// @Router			/user [post]
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
