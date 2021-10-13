package middleware

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

type RecoverErrorHandler func(ctx iris.Context, err interface{})

var (
	MaxLength    int                 = 1024 * 32
	ErrorHandler RecoverErrorHandler = func(ctx iris.Context, err interface{}) {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
	}
)

func Debug(ctx iris.Context) {
	start := time.Now()
	logger := ctx.Application().Logger()

	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			if ErrorHandler != nil {
				ErrorHandler(ctx, err)
			}
			ctx.StopExecution()

			end := time.Now()
			logger.Infof("请求协议[%v]路由[%v]，IP[%v]，Method[%v]，ContentType[%v]，返回：StatusCode[%v]，Body[%v]，执行耗时：[%v]ms\n", ctx.Request().Proto, ctx.Request().URL, ctx.RemoteAddr(), ctx.Method(), ctx.GetContentType(), ctx.GetStatusCode(), string(ctx.Recorder().Body()), end.Sub(start).Milliseconds())
		}
	}()

	rBody, err := ctx.GetBody()
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	ctx.Record()
	ctx.Next()
	end := time.Now()

	l := len(rBody)
	if l > 0 && l < MaxLength {
		logger.Infof("请求协议[%v]路由[%v]，IP[%v]，Method[%v]，ContentType[%v]，Body[%v]，返回：StatusCode[%v]，Body[%v]，执行耗时：[%v]ms\n", ctx.Request().Proto, ctx.Request().URL, ctx.RemoteAddr(), ctx.Method(), ctx.GetContentType(), string(rBody), ctx.GetStatusCode(), string(ctx.Recorder().Body()), end.Sub(start).Milliseconds())
	} else {
		logger.Infof("请求协议[%v]路由[%v]，IP[%v]，Method[%v]，ContentType[%v]，返回：StatusCode[%v]，Body[%v]，执行耗时：[%v]ms\n", ctx.Request().Proto, ctx.Request().URL, ctx.RemoteAddr(), ctx.Method(), ctx.GetContentType(), ctx.GetStatusCode(), string(ctx.Recorder().Body()), end.Sub(start).Milliseconds())
	}
}
