## Routing

### 路由行为

```cassandraql
//
app.Run(iris.Addr(":8080"), iris.WithoutPathCorrection)

///api/user and /api/user/ 表现一致
app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection)

```

### API

#### Handle方法

```cassandraql
app := iris.New()
app.Handle("GET", "/contact", func(ctx iris.Context) {
    ctx.HTML("<h1> Hello from /contact </h1>")
})
```

#### HTTP Methods (Get Post Put Delete Options Trace Connect Head Patch Any)
```cassandraql
// Method: "GET"
app.Get("/", handler)

// Method: "POST"
app.Post("/", handler)

// Method: "PUT"
app.Put("/", handler)

func handler(ctx iris.Context){
    ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}
```

### Offline Routes

> 对外隐藏，可以从其他路由处理程序中调用这个隐藏路由

```cassandraql
package main
import (
    "github.com/kataras/iris/v12"
)
func main() {
    app := iris.New()
    //直接访问/invisible/{username} 返回404
    none := app.None("/invisible/{username}", func(ctx iris.Context) {
        ctx.Writef("Hello %s with method: %s", ctx.Params().Get("username"), ctx.Method())

        if from := ctx.Values().GetString("from"); from != "" {
            ctx.Writef("\nI see that you're coming from %s", from)
        }
    })

    app.Get("/change", func(ctx iris.Context) {

        if none.IsOnline() {
            none.Method = iris.MethodNone
        } else {
            none.Method = iris.MethodGet
        }

        // refresh re-builds the router at serve-time in order to
        // be notified for its new routes.
        app.RefreshRouter()
    })

    app.Get("/execute", func(ctx iris.Context) {
        if !none.IsOnline() {
            ctx.Values().Set("from", "/execute with offline access")
            ctx.Exec("NONE", "/invisible/iris")
            return
        }

        // same as navigating to "http://localhost:8080/invisible/iris"
        // when /change has being invoked and route state changed
        // from "offline" to "online"
        ctx.Values().Set("from", "/execute")
        // values and session can be
        // shared when calling Exec from a "foreign" context.
        // 	ctx.Exec("NONE", "/invisible/iris")
        // or after "/change":
        ctx.Exec("GET", "/invisible/iris")
    })

    app.Run(iris.Addr(":8080"))
}
```

