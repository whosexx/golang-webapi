package v1

import (
	"webapi/utils"

	"github.com/kataras/iris/v12/mvc"
)

var ApiV1Routers map[string]interface{} = make(map[string]interface{}, 10)

type ApiV1Controller struct {
	utils.DependencyObject
}

//all router
func HandleRouterV1(ctx *mvc.Application, dependencies ...interface{}) *mvc.Application {
	ctx.Register(dependencies...)

	ctx.Handle(new(ApiV1Controller))
	for k, v := range ApiV1Routers {
		ctx.Party(k).Handle(v)
	}
	return ctx
}

func (api *ApiV1Controller) Get() *utils.ResultInfo {
	return &utils.ResultInfo{
		Code:    0,
		Message: "test api v1",
	}
}
