package route

import (
	v1 "golang-webapi/controllers/v1"
	v2 "golang-webapi/controllers/v2"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Handle(app *iris.Application, dependencies ...interface{}) {
	v1.HandleRouterV1(mvc.New(app.Party("/api/v1")), dependencies...)
	v2.HandleRouterV2(mvc.New(app.Party("/api/v2")), dependencies...)
}
