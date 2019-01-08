package main

import (
	"github.com/kataras/iris"
	"golang_iris_xorm/config"
	"golang_iris_xorm/route"
)

func main() {
	app := iris.New()

	// 路由初始化
	route.RouteInit(app)

	app.Run(iris.Addr("api.iris.io:8080"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:				   config.DisableStartupLog,
		TimeFormat:                        config.TimeFormat,
		Charset:                           config.Charset,
	}))
}