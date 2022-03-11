package service

import (
	"demo/app/domain"
	"demo/app/repository"
	"demo/app/serializers"
	"demo/app/utils/methods"
	"demo/infra/errors"
)

type IUsers interface {
	GetUsers() (*domain.Users, *errors.RestErr)
	CreateUser(domain.User) (*domain.User, *errors.RestErr)
	UpdateUser(userID int, req serializers.UserRequest) (*domain.User, *errors.RestErr)
	DeleteUser(userID int) *errors.RestErr
	GetUserByID(userID int) (*domain.User, *errors.RestErr)
}

type users struct {
	userRepo repository.IUsers
}

func NewUsersService(userRepo repository.IUsers) IUsers {
	return &users{
		userRepo: userRepo,
	}
}

func (u *users) GetUsers() (*domain.Users, *errors.RestErr) {
	resp, getErr := u.userRepo.All()
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}

func (u *users) CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	resp, saveErr := u.userRepo.Create(&user)
	if saveErr != nil {
		return nil, saveErr
	}
	return resp, nil
}

func (u *users) UpdateUser(userID int, req serializers.UserRequest) (*domain.User, *errors.RestErr) {
	var user *domain.User
	err := methods.StructToStruct(req, &user)
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	resp, updateErr := u.userRepo.Update(userID, user)
	if updateErr != nil {
		return nil, updateErr
	}
	return resp, nil
}

func (u *users) DeleteUser(userID int) *errors.RestErr {
	getErr := u.userRepo.Delete(userID)
	if getErr != nil {
		return getErr
	}
	return nil
}

func (u *users) GetUserByID(userID int) (*domain.User, *errors.RestErr) {
	resp, getErr := u.userRepo.Find(userID)
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}
