package v1

import (
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

var ApiV1Routers map[string]interface{} = make(map[string]interface{}, 10)

type ApiV1Controller struct {
	//	utils.DependencyObject
}

//all router
func HandleRouterV1(app *mvc.Application, dependencies ...interface{}) *mvc.Application {
	app.Register(dependencies...)

	app.Handle(new(ApiV1Controller))
	for k, v := range ApiV1Routers {
		app.Party(k).Handle(v)
	}
	return app
}

func (api *ApiV1Controller) Get() *utils.ResultInfo {
	return &utils.ResultInfo{
		Code:    0,
		Message: "test get api v1",
	}
}
