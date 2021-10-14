package services

import (
	"golang-webapi/model"
	"golang-webapi/repositories"

	"github.com/kataras/golog"
	"gorm.io/gorm"
)

type UserRepository interface {
	repositories.DBRepository

	GetAllUsers() []model.UserInfo

	GetUser(uuid string) *model.UserInfo
	GetUserByUserId(id string) *model.UserInfo
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

func (u *userService) GetAllUsers() []model.UserInfo {
	var users []model.UserInfo
	if err := u.SelectMany(&users, nil); err != nil {
		golog.Error(err.Error())
		return nil
	}
	return users
}

func (u *userService) GetUser(uuid string) *model.UserInfo {
	var user model.UserInfo
	if err := u.Select(&user, "uuid = ?", uuid); err != nil {
		golog.Error(err.Error())
		return nil
	}
	return &user
}

func (u *userService) GetUserByUserId(id string) *model.UserInfo {
	var user model.UserInfo
	if err := u.Select(&user, "userId = ?", id); err != nil {
		golog.Error(err.Error())
		return nil
	}
	return &user
}

// func (u *userService) Delete(dest interface{}, args ...interface{}) error {
// 	t := []interface{}{"uuid = ?"}
// 	t = append(t, args...)

// 	if err := u.DBService.Delete(dest, t...); err != nil {
// 		golog.Error(err.Error())
// 		return err
// 	}

// 	return nil
// }
