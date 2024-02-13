package service

import (
	"github.com/stretchr/testify/assert"
	"rest-api/model"
	"rest-api/repository/mock"
	"testing"
)

var (
	userRepositoryMock = &mock.UserRepositoryMock{}
	userService        = New(userRepositoryMock)
	user               = model.User{ID: 1, Username: "username", Firstname: "firstname", Lastname: "lastname", Age: 15}
)

func TestUserServiceFindById(t *testing.T) {
	result, _ := userService.FindById(1)
	assert.NotNil(t, result)
}

func TestUserServiceFindAll(t *testing.T) {
	result, _ := userService.FindAll()
	assert.NotNil(t, result)
}

func TestUserServiceSave(t *testing.T) {
	result, _ := userService.Save(user)
	assert.NotNil(t, result)
}
