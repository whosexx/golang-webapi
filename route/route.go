package route

import (
	v1 "golang-webapi/controllers/v1"
	v2 "golang-webapi/controllers/v2"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Handle(app *iris.Application, dependencies ...interface{}) {
	mApp := mvc.New(app.Party("/"))
	mApp.Register(dependencies...)
	mApp.HandleError(func(ctx iris.Context, err error) {
		// if !utils.IsOK(err) {
		// 	ctx.Application().Logger().Error(err)
		// }

		ctx.JSON(err)
	})

	v1.HandleRouterV1(mApp.Party("/api/v1"))
	v2.HandleRouterV2(mApp.Party("/api/v2"))
}
