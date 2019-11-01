package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/my"), myMvc)
	app.Run(iris.Addr(":8080"))
}

func myMvc(app *mvc.Application) {
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET")
}
