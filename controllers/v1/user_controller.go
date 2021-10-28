package v1

import (
	"golang-webapi/model"
	"golang-webapi/services"
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	utils.DependencyObject
	UserService services.UserRepository
}

func (c *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/all", "GetAllUsers")
	a.Handle("GET", "/info/{userId}", "GetUserByUserId")

	a.Handle("GET", "/get/{uuid}", "Get")
	a.Handle("GET", "/delete/{uuid}", "Delete")
	a.Handle("POST", "/insert", "Post")
	a.Handle("POST", "/update", "Update")
}

func (user *UserController) Get(uuid string) *utils.ResultInfo {

	u, err := user.UserService.GetUser(uuid)
	if err != nil {
		return utils.Error(err)
	}

	return utils.OK2(u)
	// return utils.ServeErr
	// panic(utils.NotFoundErr)
	// return &utils.ResultInfo{
	// 	Code:    0,
	// 	Message: "user controller get.",
	// }
}

func (user *UserController) Post() *utils.ResultInfo {
	u := &model.UserInfo{}
	if err := user.HttpContext.ReadJSON(u); err != nil {
		return utils.ServeErr.WithMessage(err.Error())
	}

	if err := user.UserService.Insert(u); err != nil {
		return utils.ServeErr.WithMessage(err.Error())
	}

	return utils.OK()
}

func (user *UserController) GetUserByUserId(userId string) *utils.ResultInfo {
	u, err := user.UserService.GetUserByUserId(userId)
	if err != nil {
		return utils.Error(err)
	}

	return utils.OK2(u)
}

// @Summary 获取用户列表1
// @Description 获取用户列表2
// @Tags　学生
// @Accept application/json
// @Produce application/json
// @Success 200 {object} utils.ResultInfo
// @Failure 200 {object} utils.ResultInfo
// @Router /api/v1/user/all [get]
func (user *UserController) GetAllUsers() *utils.ResultInfo {
	u, err := user.UserService.GetAllUsers()
	if err != nil {
		return utils.Error(err)
	}

	return utils.OK2(u)
}

func (user *UserController) Update() *utils.ResultInfo {
	u := &model.UserInfo{}
	if err := user.HttpContext.ReadJSON(u); err != nil {
		return utils.ServeErr.WithMessage(err.Error())
	}

	if err := user.UserService.Update(u); err != nil {
		return utils.ServeErr.WithMessage(err.Error())
	}

	return utils.OK()
}

func (user *UserController) Delete(uuid string) *utils.ResultInfo {
	u := &model.UserInfo{UUID: uuid}
	if err := user.UserService.Delete(u); err != nil {
		return utils.ServeErr.WithMessage(err.Error())
	}

	return utils.OK()
}
