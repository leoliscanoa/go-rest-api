package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest-api/model"
	"rest-api/service/mock"
	"testing"
)

var (
	userServiceMock     = &mock.UserServiceMock{}
	userController      = New(userServiceMock)
	user                = model.User{ID: 1, Username: "username", Firstname: "firstname", Lastname: "lastname", Age: 15}
	userString, _       = json.Marshal(user)
	requestFindById, _  = http.NewRequest(http.MethodGet, "/user/1", nil)
	requestFindAll, _   = http.NewRequest(http.MethodGet, "/user", nil)
	requestSave, _      = http.NewRequest(http.MethodPost, "/user", bytes.NewReader(userString))
	requestSaveError, _ = http.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(nil)))
	writer              = httptest.NewRecorder()
	GenericError        = errors.New("generic error")
)

func TestUserControllerFindById(t *testing.T) {
	t.Cleanup(resetState)
	vars := map[string]string{
		"id": "1",
	}
	request := mux.SetURLVars(requestFindById, vars)
	userServiceMock.On("FindById", 1).Return(model.User{}, nil)
	userController.FindById(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUserControllerFindByIdErrorId(t *testing.T) {
	t.Cleanup(resetState)
	vars := map[string]string{
		"id": "abcd",
	}
	request := mux.SetURLVars(requestFindById, vars)
	userController.FindById(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestUserControllerFindByIdError(t *testing.T) {
	t.Cleanup(resetState)
	vars := map[string]string{
		"id": "1",
	}
	request := mux.SetURLVars(requestFindById, vars)
	userServiceMock.On("FindById", 1).Return(model.User{}, GenericError)
	userController.FindById(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestUserControllerFindAll(t *testing.T) {
	t.Cleanup(resetState)
	userServiceMock.On("FindAll").Return([]model.User{}, nil)
	userController.FindAll(writer, requestFindAll)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUserControllerFindAllError(t *testing.T) {
	t.Cleanup(resetState)
	userServiceMock.On("FindAll").Return([]model.User{}, GenericError)
	userController.FindAll(writer, requestFindAll)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestUserControllerSave(t *testing.T) {
	t.Cleanup(resetState)
	requestSave.Header.Set("Content-Type", "application/json")
	userServiceMock.On("Save", user).Return(model.User{}, nil)
	userController.Save(writer, requestSave)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestUserControllerSaveErrorBody(t *testing.T) {
	t.Cleanup(resetState)
	requestSaveError.Header.Set("Content-Type", "application/json")
	userServiceMock.On("Save", user).Return(model.User{}, nil)
	userController.Save(writer, requestSaveError)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestUserControllerSaveError(t *testing.T) {
	t.Cleanup(resetState)
	requestSave.Header.Set("Content-Type", "application/json")
	userServiceMock.On("Save", user).Return(user, GenericError)
	userController.Save(writer, requestSave)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func resetState() {
	userServiceMock = &mock.UserServiceMock{}
	userController = New(userServiceMock)
	writer = httptest.NewRecorder()
	requestSave, _ = http.NewRequest(http.MethodPost, "/user", bytes.NewReader(userString))
}
