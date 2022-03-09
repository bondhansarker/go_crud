package repository

import (
	"demo/app/domain"
	"demo/infra/errors"

	"gorm.io/gorm"
)

type IUsers interface {
	Save(user *domain.User) (*domain.User, *errors.RestErr)
	All() (domain.Users, *errors.RestErr)
}

type users struct {
	*gorm.DB
}

func NewUsersRepository(db *gorm.DB) IUsers {
	return &users{
		DB: db,
	}
}

func (r *users) Save(user *domain.User) (*domain.User, *errors.RestErr) {
	res := r.DB.Model(&domain.User{}).Create(&user)

	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return user, nil
}

func (r *users) All() (domain.Users, *errors.RestErr) {
	var users domain.Users
	res := r.DB.Find(&users)

	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}

	return users, nil
}
