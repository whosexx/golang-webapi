package main

import (
	"flag"
	"fmt"
	"golang-webapi/conf"
	"golang-webapi/dependency"
	"golang-webapi/middleware"
	"golang-webapi/route"
	"golang-webapi/utils"

	"github.com/kataras/iris/v12"
)

func main() {
	t := flag.CommandLine.String("conf", "json", "conf type: json yaml toml")
	flag.Parse()

	cfg := conf.ReadConf(conf.ParseConfType(*t))

	app := iris.New()
	app.Logger().SetLevel(cfg.Level)
	app.Configure(iris.WithConfiguration(cfg.Configuration))

	app.UseGlobal(middleware.Debug)
	app.OnAnyErrorCode(func(ctx iris.Context) {
		if ctx.GetStatusCode() == iris.StatusNotFound {
			ctx.JSON(utils.NotFoundErr.WithMessage(fmt.Sprintf("route[%v] not found", ctx.Path())))
		} else {
			ctx.JSON(utils.ServeErr.WithMessage(fmt.Sprintf("invoke route[%v] err", ctx.Path())))
		}
	})
	middleware.RecoverHandler = func(ctx iris.Context, err interface{}) {
		ctx.Application().Logger().Warn(fmt.Sprintf("Recovered from a route's Handler('%s'), Exception: %v", ctx.RouteName(), err))

		if ex, ok := err.(*utils.BusinessException); !ok {
			ctx.JSON(utils.ServeErr.WithMessage(fmt.Sprintf("%s: %v", utils.ServeErr.Message, err)))
		} else {
			ctx.JSON(ex)
		}
	}

	route.Handle(app, dependency.GetDependencies(cfg)...)
	app.Listen(
		":"+fmt.Sprintf("%v", cfg.Port),
		iris.WithOptimizations,
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	//app.NewHost(&http.Server{Addr: ":8080"}).ListenAndServe()
}
