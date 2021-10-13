package v1

import (
	"golang-webapi/model"
	"golang-webapi/services"
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	ApiV1Routers["/user"] = new(UserController)
}

type UserController struct {
	utils.DependencyObject
	UserService services.UserRepository
}

func (c *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/all", "GetAllUsers")
	a.Handle("GET", "/info/{userId}", "GetUserByUserId")

	a.Handle("GET", "/", "Get")
	a.Handle("POST", "/", "Post")
}

func (user *UserController) Get() *utils.ResultInfo {
	panic(utils.NotFoundErr)
	return &utils.ResultInfo{
		Code:    0,
		Message: "user controller get.",
	}
}

func (user *UserController) Post() *utils.ResultInfo {
	u := &model.UserInfo{}
	if err := user.HttpContext.ReadJSON(u); err != nil {
		return &utils.ResultInfo{
			Code:    1000,
			Message: err.Error(),
		}
	}

	if err := user.UserService.Insert(u); err != nil {
		return &utils.ResultInfo{
			Code:    1001,
			Message: err.Error(),
		}
	}

	return &utils.ResultInfo{
		Code:    0,
		Message: "ok.",
	}
}

func (user *UserController) GetUserByUserId(userId string) *utils.ResultInfo {
	u := user.UserService.GetUserByUserId(userId)
	if u == nil {
		return utils.NewResultInfo(utils.NotFoundErr)
	}

	return &utils.ResultInfo{
		Code:    0,
		Message: "ok",
		Data:    u,
	}
}

func (user *UserController) GetAllUsers() *utils.ResultInfo {
	return &utils.ResultInfo{
		Code:    0,
		Message: "ok",
		Data:    user.UserService.GetAllUsers(),
	}
}
