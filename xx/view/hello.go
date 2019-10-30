package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	tmpl := iris.HTML("./tpl", ".html")
	tmpl.Reload(true)

	app.RegisterView(tmpl)

	app.Get("/hi", func(ctx iris.Context) {
		ctx.ViewData("message", "hello iris view!")
		ctx.View("hi.html")
	})

	app.Run(iris.Addr(":8080"))
}
