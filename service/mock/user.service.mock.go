package mock

import (
	"github.com/stretchr/testify/mock"
	"rest-api/model"
)

type UserServiceMock struct{ mock.Mock }

func (m *UserServiceMock) FindById(id int) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
func (m *UserServiceMock) FindAll() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}
func (m *UserServiceMock) Save(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}
