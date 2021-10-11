package main

import (
	"flag"
	"fmt"
	"golang-webapi/conf"
	"golang-webapi/dependency"
	"golang-webapi/route"

	"github.com/kataras/iris/v12"
)

func main() {
	t := flag.CommandLine.String("conf", "json", "conf type: json yaml toml")
	flag.Parse()

	cfg := conf.ReadConf(conf.ParseConfType(*t))

	app := iris.New()
	app.Logger().SetLevel(cfg.Level)
	app.Configure(iris.WithConfiguration(cfg.Configuration))

	route.Handle(app, dependency.GetDependencies(cfg)...)
	app.Listen(":" + fmt.Sprintf("%v", cfg.Port))
}
