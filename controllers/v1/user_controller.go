package v1

import (
	"fmt"
	"golang-webapi/model"
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	ApiV1Routers["/user"] = new(UserController)
}

type UserController struct {
	utils.DependencyObject
}

func (c *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/info/{userId}", "GetInfoByUserId")
}

func (user *UserController) Get() *utils.ResultInfo {
	return &utils.ResultInfo{
		Code:    0,
		Message: "user controller get.",
	}
}

func (c *UserController) GetInfoByUserId(userId string) *utils.ResultInfo {
	var users []model.UserInfo
	c.DBContext.Where("userId = ?", userId).Find(&users)
	fmt.Println(users)

	return &utils.ResultInfo{
		Code:    0,
		Message: "ok",
		Data:    &users,
	}
}
