package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))
	return app
}

type ExampleController struct{}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "hello iris!"}
}

func (c *ExampleController) GetUserName() interface{} {
	return map[string]string{"username": "limq!"}
}

func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	//anyMiddlewareHere := func(ctx iris.Context){
	//	ctx.Application().Logger().Warnf("Inside / custom_path")
	//	ctx.Next()
	//}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
}

func (c ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from custom handler without following the naming guide"
}

func anyMiddlewareHere(ctx iris.Context) {
	ctx.Application().Logger().Warnf("Inside / custom_path")
	ctx.Next()
}
