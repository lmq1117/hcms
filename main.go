package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/basic"), basicMvc)

	app.Run(iris.Addr(":8080"))
}

func basicMvc(app *mvc.Application) {
	//使用中间件
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path:%s", ctx.Path())
		ctx.Next()
	})
	//
	app.Register(
		sessions.New(sessions.Config{}).Start,
	)

	//Handle控制器
	app.Handle(new(basicController))
	app.Party("/sub").Handle(new(basicSubController))

}
