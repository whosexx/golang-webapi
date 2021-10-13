package utils

import (
	"golang-webapi/conf"

	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	OK          = NewBusinessException(0, "ok")
	NotFoundErr = NewBusinessException(404, "not found")
	ServeErr    = NewBusinessException(500, "serve error")
)

type BusinessException struct {
	Code    int    `json:"code"`
	Err     string `json:"error"`
	Message string `json:"message"`
}

func (ex *BusinessException) Error() string {
	return ex.Err
	//return fmt.Sprintf("%s(%d)", ex.Err, ex.Code)
}

func NewBusinessException(code int, err string) *BusinessException {
	return &BusinessException{
		Code: code,
		Err:  err,
	}
}

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

func NewResultInfo(ex *BusinessException, msgs ...string) *ResultInfo {
	msg := ""
	if len(msgs) > 0 {
		for _, m := range msgs {
			msg = msg + m
		}
	} else {
		msg = ex.Message
		if msg == "" {
			msg = ex.Err
		}
	}

	return &ResultInfo{
		Code:    ex.Code,
		Message: msg,
	}
}
