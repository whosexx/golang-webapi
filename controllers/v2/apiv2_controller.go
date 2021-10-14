package v2

import (
	"golang-webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

var ApiV2Routers map[string]interface{} = map[string]interface{}{}

type ApiV2Controller struct {
	utils.DependencyObject
}

//all router
func HandleRouterV2(ctx *mvc.Application, dependencies ...interface{}) *mvc.Application {
	for k, v := range ApiV2Routers {
		ctx.Party(k).Handle(v)
	}

	ctx.Handle(new(ApiV2Controller))
	return ctx
}

func (api *ApiV2Controller) Get() *utils.ResultInfo {
	return utils.OK("test api v2")
}
