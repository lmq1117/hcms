package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

	//app.SetExecutionRules(iris.ExecutionRules{
	//	Begin:iris.ExecutionOptions{Force:true},
	//	Main:iris.ExecutionOptions{Force:true},
	//	Done:iris.ExecutionOptions{Force:true},
	//})

	app.UseGlobal(globalBefore)
	app.DoneGlobal(globalAfter)

	app.Run(iris.Addr(":8080"))
}

func globalBefore(ctx iris.Context) {
	println("path:" + ctx.Path() + " globalBefore")
	ctx.Next()
}

func globalAfter(ctx iris.Context) {
	println("path:" + ctx.Path() + " globalAfter")
	println(`before the party's routes and its children,
but this is not applied to the '/' route
because it's registered before the middleware, order matters.`)
	ctx.Next()
}

func indexHandler(ctx iris.Context) {
	ctx.HTML("<h1>Index</h1>")
	ctx.Next()
}

func contactHandler(ctx iris.Context) {
	ctx.HTML("<h1>Contact</h1>")
	ctx.Next()
}
