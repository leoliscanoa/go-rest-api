package repository

import (
	"gorm.io/gorm"
	"log"
	"rest-api/model"
)

type UserRepositoryImplement struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return &UserRepositoryImplement{db: db}
}

func (u UserRepositoryImplement) FindById(id int) (model.User, error) {
	var user = model.User{ID: uint64(id)}
	tx := u.db.First(&user)
	if tx.Error != nil {
		log.Printf("Error: %v", tx.Error.Error())
	}
	return user, tx.Error
}

func (u UserRepositoryImplement) FindAll() ([]model.User, error) {
	var users []model.User
	tx := u.db.Find(&users)
	if tx.Error != nil {
		log.Printf("Error: %v", tx.Error.Error())
	}
	return users, tx.Error
}

func (u UserRepositoryImplement) Save(user model.User) (model.User, error) {
	tx := u.db.Create(&user)
	if tx.Error != nil {
		log.Printf("Error: %v", tx.Error.Error())
	}
	return user, tx.Error
}
