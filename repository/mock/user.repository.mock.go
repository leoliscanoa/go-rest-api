package mock

import "rest-api/model"

type UserRepositoryMock struct{}

func (m *UserRepositoryMock) FindById(id int) (model.User, error) {
	return model.User{}, nil
}
func (m *UserRepositoryMock) FindAll() ([]model.User, error) {
	return []model.User{}, nil
}
func (m *UserRepositoryMock) Save(user model.User) (model.User, error) {
	return user, nil
}
