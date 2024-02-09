package service

import (
	"rest-api/model"
	"rest-api/repository"
)

type UserServiceImplement struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) UserService {
	return &UserServiceImplement{userRepository: userRepository}
}

func (u *UserServiceImplement) FindById(id int) (model.User, error) {
	return u.userRepository.FindById(id)
}

func (u *UserServiceImplement) FindAll() ([]model.User, error) {
	return u.userRepository.FindAll()
}

func (u *UserServiceImplement) Save(user model.User) (model.User, error) {
	return u.userRepository.Save(user)
}
