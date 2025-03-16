package v2

import (
	"errors"
	"github.com/vlbarou/sampleproject/internal/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}
