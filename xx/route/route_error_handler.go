package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>success Index</h1>")
}

func notFound(ctx iris.Context) {
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("服务器内部错误 请再次尝试")
}
