package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	mock2 "rest-api/config/mock"
	"rest-api/model"
	"testing"
)

var (
	userRow = sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15)
	usersRows = sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
			AddRow(1, "username", "firstname", "lastname", 15).
			AddRow(2, "username", "firstname", "lastname", 15)
	user           = model.User{ID: 1, Username: "username", Firstname: "firstname", Lastname: "lastname", Age: 15}
	GenericError   = errors.New("generic error")
	database, mock = mock2.MockDatabase()
	implMock       = New(database)
)

func TestUserRepositoryFindById(t *testing.T) {
	mock.ExpectQuery(`SELECT`).WillReturnRows(userRow)
	response, _ := implMock.FindById(1)
	assert.Equal(t, uint64(1), response.ID)
}

func TestUserRepositoryFindByIdError(t *testing.T) {
	mock.ExpectQuery(`SELECT`).WillReturnError(GenericError)
	_, err := implMock.FindById(1)
	assert.Equal(t, err, GenericError)
}

func TestUserRepositoryFindAll(t *testing.T) {
	mock.ExpectQuery(`SELECT`).WillReturnRows(usersRows)
	response, _ := implMock.FindAll()
	assert.Equal(t, 2, len(response))
}

func TestUserRepositoryFindAllError(t *testing.T) {
	mock.ExpectQuery(`SELECT`).WillReturnError(GenericError)
	_, err := implMock.FindAll()
	assert.Equal(t, err, GenericError)
}

func TestUserRepositorySave(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT`).
		WithArgs("username", "firstname", "lastname", 15, 1).
		WillReturnRows(userRow)
	mock.ExpectCommit()
	result, _ := implMock.Save(user)
	assert.Equal(t, user.ID, result.ID)
}

func TestUserRepositorySaveError(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT`).
		WithArgs("username", "firstname", "lastname", 15, 1).
		WillReturnError(GenericError)
	mock.ExpectRollback()
	_, err := implMock.Save(user)
	assert.Equal(t, err, GenericError)
}
