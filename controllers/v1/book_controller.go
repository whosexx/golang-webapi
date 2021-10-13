package v1

import (
	"golang-webapi/utils"

	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12/mvc"
)

//register url
func init() {
	ApiV1Routers["/book"] = new(BookController)
}

type BookController struct {
	utils.DependencyObject
}

func (book *BookController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "Get")
	app.Handle("GET", "/redis", "GetRedis")
}

func (book *BookController) Get() *utils.ResultInfo {
	return utils.OK()
}

func (book *BookController) GetRedis() *utils.ResultInfo {
	conn := book.Redis.Get()
	defer conn.Close()

	d, _ := redis.String(conn.Do("info"))
	book.Logger.Debug(d)
	return utils.OK()
}
