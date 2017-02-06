# Baa 核心

## 创建应用

`func New() *Baa`

快速创建一个新的应用实例。

`func Instance(name string) *Baa`

获取一个命名实例，如果实例不存在则调用 `New()` 创建并且命名。

命名实例 用于不同模块间共享 baa实例的场景。在 入口中创建，在其他模块中 `baa.Instance(name)` 可获取指定的实例。

`func Default() *Baa`

使用默认的应用实例。

Default 是 `Instance` 的一个默认实现，是全局唯一的实例。在共享场景下，不需要传递 baa 直接调用 `baa.Default()` 即可访问同一实例。

使用示例：

```
app := baa.New()
app := baa.Default()
app := baa.Instance("myApp")
myApp := baa.Instance("myApp")
```

## 路由管理

baa 基于 http resetfull 模式设计了路由管理器，具体见 [路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md) 一节。

## 中间件

baa 支持通过 中间件 机制，注入请求过程，实现类似插件的功能，具体见 [中间件](https://github.com/go-baa/doc/tree/master/zh-CN/middleware.md) 一节。

## 依赖注入

依赖注入(dependency injection)简称 DI，是 baa 实现的核心，baa 所有组件基于DI组装起来的。

baa组件的更换，见 [更换内置引擎](#更换内置引擎) 一节。

DI的具体使用，见 [依赖注入](https://github.com/go-baa/doc/tree/master/zh-CN/di.md) 一节。

## 运行应用

`func (b *Baa) Run(addr string)`

指定一个监听地址，启动一个HTTP服务。

示例：

```
app := baa.Default()
app.Run(":1323")
```

`func (b *Baa) RunTLS(addr, certfile, keyfile string)`

指定监听地址和TLS证书，启动一个HTTPS服务。

示例：

```
app := baa.Default()
app.RunTLS(":8443", "cert/cert.cert", "cert/server.key")
```

## 环境变量

`BAA_ENV`

baa 通过 系统环境变量 `BAA_ENV` 来设置运行模式。

`baa.Env` 

外部程序可以通过 `baa.Env` 变量来获取 baa 当前的运行模式。

运行模式常量

```
// DEV mode
DEV = "development"
// PROD mode
PROD = "production"
// TEST mode
TEST = "test"
```

* baa.DEV  开发模式
* baa.PROD 产品模式
* baa.TEST 测试模式

示例代码：

```
if baa.Env == baa.PROD {
    // 应用运行在产品模式
}
```

## 调试

`func (b *Baa) Debug() bool`

返回是否是调试模式，应用可以根据是否运行在调试模式，来输出调试信息。

`func (b *Baa) SetDebug(v bool)`

默认根据运行环境决定是否开启调试模式，可以通过该方法开启/关闭调试模式。

> 在 产品模式 下，默认关闭调试模式，其他模式下默认开启调试模式。

`func (b *Baa) Logger() Logger`

返回日志器，在应用中可以调用日志器来输出日志。

示例：

```
app := baa.New()
log := app.Logger()
log.Println("test")
```

## 错误处理

> 错误输出，只是给浏览器返回错误，但并不会阻止接下来绑定的方法。

`func (b *Baa) NotFound(c *Context)`

调用该方法会直接 输出 404错误。

`func (b *Baa) Error(err error, c *Context)`

调用该方法会直接输出 500错误，并根据运行模式决定是否在浏览器中返回具体错误。

示例

```
app := baa.New()
app.Get("/", func(c *baa.Context) {
    c.Baa().NotFound(c)
})
app.Get("/e", func(c *baa.Context) {
    c.Baa().Error(errors.New("something error"), c)
})

```

## 更换内置引擎

baa 采用以DI为核心的框架设计，内置模块均可使用新的实现通过DI更换。

### 日志器

baa 将日志抽象为 `baa.Logger` 接口，只要实现了该接口，就可以注册为日志器。

baa 内置的日志器使用的是标准包的 `log` 实例。

更换日志器：

```
app := baa.New()
baa.SetDI("logger", newLogger)
```

> logger 是内置名称，该命名被用于全局日志器。

### 路由器

只要实现接口 `baa.Router` 接口即可。

```
app := baa.New()
baa.SetDI("router", newRouter)
```

> router 是内置名称，该命名被用于全局路由器。

baa 除了内置的 tree路由，还新增了两个路由器可用，见 [router](https://github.com/go-baa/router)

### 模板引擎

只要实现接口 `baa.Renderer` 接口即可。

```
app := baa.New()
baa.SetDI("render", newRender)
```

> render 是内置名称，该命名被用于模板渲染。

### DIer

甚至依赖注入管理器，自己也能被替换，只要实现 `baa.Dier` 接口即可。

请注意要在第一个设置，并且重设以上三个引擎，因为你的注入管理器中默认并没有内置引擎，BAA将发生错误。

```
app := baa.New()
app.SetDIer(newDIer)
app.SetDI("logger", log.New(os.Stderr, "[Baa] ", log.LstdFlags))
app.SetDI("render", new(baa.Render))
app.SetDI("router", baa.NewTree(app))
```
