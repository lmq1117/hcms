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
> 1 直接访问 http://localhost:8080/invisible/iris 404
> 2 访问 http://localhost:8080/execute 执行到none
> 3 访问 http://localhost:8080/change 后 可以直接访问 http://localhost:8080/invisible/iris
### 路由分组
#### .Party()
```cassandraql
app := iris.New()

users := app.Party("/users",myAuthMiddlewareHandler)
users.Get("/{id:uint64}/profile",userProfileHandler)
users.Get("/message/{id:uint64}",userMessageHandler)
```

#### .PartyFunc()
```cassandraql
app := iris.New()
app.PartyFunc("/users", func(users iris.Party) {
    users.Use(myAuthMiddlewareHandler)

    // http://localhost:8080/users/42/profile
    users.Get("/{id:uint64}/profile", userProfileHandler)
    // http://localhost:8080/users/messages/1
    users.Get("/messages/{id:uint64}", userMessageHandler)
})
```

### Path Parameters 
> ctx.Params() 获取路由中定义的url路径变量 ?? 与ctx.Values()区别??
>> path中的路由参数可以从ctx.Params()中取得
>> 被用于中间件和处理程序交互的 ctx.Values().???
```cassandraql
app.Get("/u/{username:string}", func(ctx iris.Context) {
	ctx.Writef("username (string): %s", ctx.Params().Get("username"))
})
```

### 中间件
#### 例子
#### 全局中间件
> 使用ExecutionRules强制使用handle而用写ctx.Next()
#### 内置中间件
####  社区中间件

### 错误管理400 500
>OnErrorCode

>Context.Problem json xml

### 子域名 先跳过2019-11-01

## mvc

### hello world

#### 常规用法
```cassandraql
app := iris.New()
mvc.New(app).Handle(new(ExampleController))
app.Run(iris.Addr(":8080"))
```

#### 默认命名规则及对应path
> Get：get /

> GetPing:get /ping

> GetUserName:get /user/name

#### 不遵循默认命名规则

```cassandraql
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
}
```
### basic





