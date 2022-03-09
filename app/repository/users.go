package repository

import (
	"demo/app/domain"
	"demo/infra/errors"
	"gorm.io/gorm"
)

type IUsers interface {
	All() (*domain.Users, *errors.RestErr)
	Create(user *domain.User) (*domain.User, *errors.RestErr)
	Update(userID int, user *domain.User) (*domain.User, *errors.RestErr)
	Find(userID int) (*domain.User, *errors.RestErr)
	Delete(userID int) *errors.RestErr
}

type users struct {
	*gorm.DB
}

func NewUsersRepository(db *gorm.DB) IUsers {
	return &users{
		DB: db,
	}
}

func (r *users) All() (*domain.Users, *errors.RestErr) {
	var users *domain.Users
	res := r.DB.Find(&users)
	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	return users, nil
}

func (r *users) Create(user *domain.User) (*domain.User, *errors.RestErr) {
	response := r.DB.Model(&domain.User{}).Create(&user)
	if response.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	return user, nil
}

func (r *users) Update(userID int, user *domain.User) (*domain.User, *errors.RestErr) {
	res := r.DB.Model(&domain.User{}).Where("id = ?", userID).Updates(&user)
	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	return user, nil
}

func (r *users) Find(userID int) (*domain.User, *errors.RestErr) {
	var user *domain.User
	res := r.DB.Model(&domain.User{}).Where("id = ?", userID).First(&user)
	if res.RowsAffected == 0 {
		return nil, errors.NewNotFoundError(errors.ErrRecordNotFound)
	}
	if res.Error != nil {
		return nil, errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	return user, nil
}

func (r *users) Delete(userID int) *errors.RestErr {
	res := r.DB.Delete(&domain.User{}, userID)
	if res.RowsAffected == 0 {
		return errors.NewNotFoundError(errors.ErrRecordNotFound)
	}
	if res.Error != nil {
		return errors.NewInternalServerError(errors.ErrSomethingWentWrong)
	}
	return nil
}
