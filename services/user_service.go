package services

import (
	"golang-webapi/model"

	"github.com/kataras/golog"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() []model.UserInfo
	GetUserByUserId(id string) *model.UserInfo

	Insert(user interface{}) error
	Update(user interface{}) error
	Delete(id string) error
}

type userService struct {
	DataContext *gorm.DB
}

func NewUserService(db *gorm.DB) UserRepository {
	return &userService{
		DataContext: db,
	}
}

func (u *userService) GetAllUsers() []model.UserInfo {
	var users []model.UserInfo
	if err := u.DataContext.Find(&users).Error; err != nil {
		golog.Error(err.Error())
		return nil
	}
	return users
}

func (u *userService) GetUserByUserId(id string) *model.UserInfo {
	var user model.UserInfo
	if err := u.DataContext.Where("userId = ?", id).First(&user).Error; err != nil {
		golog.Error(err.Error())
		return nil
	}
	return &user
}

func (u *userService) Insert(user interface{}) error {
	if err := u.DataContext.Create(user).Error; err != nil {
		golog.Error(err.Error())
		return err
	}
	return nil
}

func (u *userService) Update(user interface{}) error {
	if err := u.DataContext.Save(user).Error; err != nil {
		golog.Error(err.Error())
		return err
	}
	return nil
}

func (u *userService) Delete(id string) error {
	var user model.UserInfo
	if err := u.DataContext.Where("userId = ?", id).First(&user).Error; err != nil {
		golog.Error(err.Error())
		return err
	}

	u.DataContext.Delete(&user)
	return nil
}
