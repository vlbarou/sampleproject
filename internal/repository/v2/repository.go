package v2

import "github.com/vlbarou/sampleproject/internal/model"

// https://chatgpt.com/share/67d612e5-21b4-8009-9c5f-f02fec6e375b

type UserRepository interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
}
