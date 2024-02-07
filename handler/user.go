package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/constants"
	"rest-api/dto"
	"rest-api/model"
	"strconv"
	"time"
)

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	db := constants.Database
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		var response = dto.ResponseDto[string]{
			Message: "Error",
			Time:    time.Now().UTC().String(),
			Error:   err.Error(),
		}
		json.NewEncoder(writer).Encode(response)
	}
	var user = model.User{ID: uint64(id)}
	db.First(&user)
	var response = dto.ResponseDto[model.User]{
		Message: "Usuario encontrado por id",
		Time:    time.Now().UTC().String(),
		Data:    user,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	db := constants.Database
	var users []model.User
	result := db.Find(&users)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	var response = dto.ResponseDto[[]model.User]{
		Message:       "Lista de usuarios",
		Time:          time.Now().UTC().String(),
		Data:          users,
		TotalElements: int(result.RowsAffected),
	}
	json.NewEncoder(writer).Encode(response)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	db := constants.Database
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		var response = dto.ResponseDto[string]{
			Message: "Error",
			Time:    time.Now().UTC().String(),
			Error:   err.Error(),
		}
		json.NewEncoder(writer).Encode(response)
	}
	db.Create(&user)
	var response = dto.ResponseDto[model.User]{
		Message: "Usuario creado",
		Time:    time.Now().UTC().String(),
		Data:    user,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}
