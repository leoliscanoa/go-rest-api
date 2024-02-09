package repository

import "rest-api/model"

type UserRepository interface {
	FindById(id int) (model.User, error)
	FindAll() ([]model.User, error)
	Save(user model.User) (model.User, error)
}
