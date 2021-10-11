package v1

import (
	"fmt"
	"golang-webapi/utils"

	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12/mvc"
)

func init() {
	ApiV1Routers["/book"] = new(BookController)
}

type BookController struct {
	utils.DependencyObject
}

func (book *BookController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetBookInfo")
}

func (book *BookController) GetBookInfo() *utils.ResultInfo {
	conn := book.Redis.Get()
	defer conn.Close()

	d, _ := redis.String(conn.Do("info"))
	fmt.Println(d)
	return &utils.ResultInfo{
		Code:    0,
		Message: "book",
	}
}
