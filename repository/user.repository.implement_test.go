package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestFindById(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	defer mockDb.Close()
	implMock := New(db)
	user := sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15)
	mock.ExpectQuery(`SELECT`).WillReturnRows(user)
	response, _ := implMock.FindById(1)
	assert.Equal(t, uint64(1), response.ID)
}
func TestFindByIdError(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	defer mockDb.Close()
	implMock := New(db)
	UserNotFound := errors.New("user not found")
	mock.ExpectQuery(`SELECT`).WillReturnError(UserNotFound)
	_, err := implMock.FindById(1)
	assert.Equal(t, err, UserNotFound)
}

func TestFindAll(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	defer mockDb.Close()
	implMock := New(db)
	users := sqlmock.NewRows([]string{"id", "username", "firstname", "lastname", "age"}).
		AddRow(1, "username", "firstname", "lastname", 15).
		AddRow(2, "username", "firstname", "lastname", 15)
	mock.ExpectQuery(`SELECT`).WillReturnRows(users)
	response, _ := implMock.FindAll()
	assert.Equal(t, 2, len(response))
}
