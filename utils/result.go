package utils

import (
	"golang-webapi/conf"

	"github.com/gomodule/redigo/redis"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type DependencyObject struct {
	Conf        *conf.Conf
	HttpContext iris.Context
	DBContext   *gorm.DB
	Redis       *redis.Pool
	Logger      *golog.Logger
	//Session     *sessions.Session
}

type ResultInfo struct {
	Code    int         `json:"code"`
	Err     string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (ex *ResultInfo) WithData(data interface{}) *ResultInfo {
	ex.Data = data
	return ex
}

func (ex *ResultInfo) WithMessage(msg string) *ResultInfo {
	ex.Message = msg
	return ex
}
