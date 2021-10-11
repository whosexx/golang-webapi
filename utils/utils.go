package utils

import (
	"webapi/conf"

	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type DependencyObject struct {
	Conf        *conf.Conf
	HttpContext iris.Context
	DBContext   *gorm.DB
	Redis       *redis.Pool
}

type ResultInfo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
