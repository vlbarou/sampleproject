package v2

import (
	"github.com/stretchr/testify/mock"
	"github.com/vlbarou/sampleproject/internal/model"
)

type MockUserRepository struct {
	mock.Mock
	//MockCreateUser  func(user *model.User) error
	//MockGetAllUsers func() ([]model.User, error)
	//MockGetUserByID func(id uint) (*model.User, error)
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id uint) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

//func (m *MockUserRepository) CreateUser(user *model.User) error {
//	return m.MockCreateUser(user)
//}
//
//func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
//	return m.MockGetAllUsers()
//}
//
//func (m *MockUserRepository) GetUserByID(id uint) (*model.User, error) {
//	return m.MockGetUserByID(id)
//}
