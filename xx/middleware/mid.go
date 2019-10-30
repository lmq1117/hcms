package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/", before, mainHandler, after)
	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
	shareInformation := "这是共享信息 在中间件和处理程序之间的"
	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)
	ctx.Values().Set("info", shareInformation)
	ctx.Next()
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	info := ctx.Values().GetString("info")
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)
	ctx.Next()
}
