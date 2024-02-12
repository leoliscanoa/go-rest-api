package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"rest-api/config"
	"rest-api/model"
	"testing"
)

func TestFindById(t *testing.T) {
	database, mock := config.MockDatabase()
	implMock := New(database)
	user := sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15)
	mock.ExpectQuery(`SELECT`).WillReturnRows(user)
	response, _ := implMock.FindById(1)
	assert.Equal(t, uint64(1), response.ID)
}

func TestFindByIdError(t *testing.T) {
	database, mock := config.MockDatabase()
	implMock := New(database)
	UserNotFound := errors.New("user not found")
	mock.ExpectQuery(`SELECT`).WillReturnError(UserNotFound)
	_, err := implMock.FindById(1)
	assert.Equal(t, err, UserNotFound)
}

func TestFindAll(t *testing.T) {
	database, mock := config.MockDatabase()
	implMock := New(database)
	users := sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15).
		AddRow(2, "username", "firstname", "lastname", 15)
	mock.ExpectQuery(`SELECT`).WillReturnRows(users)
	response, _ := implMock.FindAll()
	assert.Equal(t, 2, len(response))
}

func TestFindAllError(t *testing.T) {
	database, mock := config.MockDatabase()
	implMock := New(database)
	UsersNotFound := errors.New("users not found")
	mock.ExpectQuery(`SELECT`).WillReturnError(UsersNotFound)
	_, err := implMock.FindAll()
	assert.Equal(t, err, UsersNotFound)
}

func TestSave(t *testing.T) {
	database, mock := config.MockDatabase()
	implMock := New(database)
	user := model.User{ID: 1, Username: "username", Firstname: "firstname", Lastname: "lastname", Age: 15}
	userRow := sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT`).
		WithArgs("username", "firstname", "lastname", 15, 1).
		WillReturnRows(userRow)
	mock.ExpectCommit()
	result, _ := implMock.Save(user)
	assert.Equal(t, user.ID, result.ID)
}
