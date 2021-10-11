package main

import (
	"fmt"
	"webapi/conf"
	"webapi/dependency"
	"webapi/route"

	"github.com/kataras/iris/v12"
)

func main() {
	cfg := conf.ReadConf()

	app := iris.New()
	app.Logger().SetLevel(cfg.Level)
	app.Configure(iris.WithConfiguration(cfg.Configuration))

	route.Handle(app, dependency.GetDependencies(cfg)...)
	app.Listen(":" + fmt.Sprintf("%v", cfg.Port))
}
