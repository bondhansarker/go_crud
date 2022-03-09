package service

import (
	"demo/app/domain"
	"demo/app/repository"
	"demo/infra/errors"
)

type IUsers interface {
	CreateUser(domain.User) (*domain.User, *errors.RestErr)
	GetUsers() (domain.Users, *errors.RestErr)
}

type users struct {
	urepo repository.IUsers
}

func NewUsersService(urepo repository.IUsers) IUsers {
	return &users{
		urepo: urepo,
	}
}

func (u *users) CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	resp, saveErr := u.urepo.Save(&user)
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}

func (u *users) GetUsers() (domain.Users, *errors.RestErr) {
	resp, getErr := u.urepo.All()
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}
