package database

import (
	"github.com/dyhalmeida/go-apis/internal/entity"
	"gorm.io/gorm"
)

type UserDbInterface interface {
	Create(*entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type UserDb struct {
	DB *gorm.DB
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{
		DB: db,
	}
}

func (u *UserDb) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDb) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
