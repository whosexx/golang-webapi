package services

import (
	"golang-webapi/model"
	"golang-webapi/repositories"

	"gorm.io/gorm"
)

type UserRepository interface {
	repositories.DBRepository

	GetAllUsers() ([]model.UserInfo, error)

	GetUser(uuid string) (*model.UserInfo, error)
	GetUserByUserId(id string) (*model.UserInfo, error)
}

type userService struct {
	repositories.DBService
}

func NewUserService(db *gorm.DB) UserRepository {
	return &userService{
		DBService: repositories.DBService{
			DataContext: db,
		},
	}
}

func (u *userService) GetAllUsers() ([]model.UserInfo, error) {
	var users []model.UserInfo
	if err := u.SelectMany(&users, nil); err != nil {
		return users, err
	}

	return users, nil
}

func (u *userService) GetUser(uuid string) (*model.UserInfo, error) {
	var user model.UserInfo
	if err := u.Select(&user, "uuid = ?", uuid); err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *userService) GetUserByUserId(id string) (*model.UserInfo, error) {
	var user model.UserInfo
	if err := u.Select(&user, "userId = ?", id); err != nil {
		return &user, err
	}

	return &user, nil
}
