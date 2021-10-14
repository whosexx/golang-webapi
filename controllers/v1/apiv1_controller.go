package v1

import (
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

//register controller
var ApiV1Routers map[string]interface{} = map[string]interface{}{
	"/user": new(UserController),
	"/book": new(BookController),
}

type ApiV1Controller struct {
	//	utils.DependencyObject
}

//all router
func HandleRouterV1(app *mvc.Application, dependencies ...interface{}) *mvc.Application {
	app.Handle(new(ApiV1Controller))
	for k, v := range ApiV1Routers {
		app.Party(k).Handle(v)
	}
	return app
}

func (api *ApiV1Controller) Get() *utils.ResultInfo {
	return utils.OK("test get api v1")
}
