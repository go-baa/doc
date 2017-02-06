# Baa 依赖注入

依赖注入(dependency injection)简称 DI，是 baa 实现的核心，baa 所有组件基于DI组装起来的。

默认的 日志、路由、模板 都是通过 DI 注册进来的，在 [Baa核心#更换内置引擎](https://github.com/go-baa/doc/blob/master/zh-CN/baa.md)一节也介绍过。

Baa的初始化函数是这样写的：

```
// New create a baa application without any config.
func New() *Baa {
	b := new(Baa)
	b.middleware = make([]HandlerFunc, 0)
	b.pool = sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, b)
		},
	}
	if Env != PROD {
		b.debug = true
	}
	b.SetDIer(NewDI())
	b.SetDI("router", NewTree(b))
	b.SetDI("logger", log.New(os.Stderr, "[Baa] ", log.LstdFlags))
	b.SetDI("render", newRender())
	b.SetNotFound(b.DefaultNotFoundHandler)
	return b
}
```

代码出处，[baa.go](https://github.com/go-baa/baa/blob/master/baa.go)

## 注册

`func (b *Baa) SetDI(name string, h interface{})`

DI 的注册和使用都依赖于注册的名称，比如要更换内置组件必须注册为指定的名称。

### name string

依赖的名称

### h interface{}

依赖的实例，可以是任意类型

使用示例：

```
package main

import (
	"log"
	"os"

	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.SetDI("logger", log.New(os.Stderr, "[BaaDI] ", log.LstdFlags))

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

## 使用

`func (b *Baa) GetDI(name string) interface{}`

`func (c *Context) DI(name string) interface{}`


可以通过 `baa.GetDI` 或者 `c.DI` 来获取已经注册的依赖，如果未注册，返回 `nil`

由于注册的依赖可能是任意类型，故返回类型为 `interface{}`，所以获取后，需要做一次类型断言再使用。

使用示例：

```
package main

import (
	"log"
	"os"

	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.SetDI("logger", log.New(os.Stderr, "[BaaDI] ", log.LstdFlags))

	app.Get("/", func(c *baa.Context) {
		// use di
		logger := c.DI("logger").(*log.Logger)
		logger.Println("i am use logger di")

		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

### 日志

baa 将日志抽象为 `baa.Logger` 接口，只要实现了该接口，就可以注册为全局日志器。

baa 内置的日志器使用的是标准包的 `log` 实例。

更换全局日志器：

```
app := baa.New()
baa.SetDI("logger", newLogger)
```

> logger 是内置名称，该命名被用于全局日志器。
> 
> 如果不是要更换全局日志，而是注册一个新的日志器用于其他用途，只需更改注册名称即可，而且也不需要实现 `baa.Logger` 接口。

### 路由

只要实现接口 `baa.Router` 接口即可。

```
app := baa.New()
baa.SetDI("router", newRender)
```

> router 是内置名称，，该命名被用于全局路由器。

baa 除了内置的 tree路由，还新增了两个路由器可用，见 [router](https://github.com/go-baa/router)

```
package main

import (
    "github.com/go-baa/router/regtree"
    "gopkg.in/baa.v1"
)

func main() {
    app := baa.Default()
    app.SetDI("router", regtree.New(app))

    app.Get("/view-:id(\\d+)", func(c *baa.Context) {
        c.String(200, c.Param("id"))
    })
    app.Get("/view-:id(\\d+)/project", func(c *baa.Context) {
        c.String(200, c.Param("id")+"/project")
    })
    app.Run(":1323")
}
```

### 模板

只要实现接口 `baa.Renderer` 接口即可。

```
app := baa.New()
baa.SetDI("render", newRender)
```

> render 是内置名称，该命名被用于模板渲染。

baa 除了内置的 render简单模板渲染，还新增了两个模板渲染引擎：

* [render](https://github.com/go-baa/doc/tree/master/zh-CN/component/render.md)
* [pongo2](https://github.com/go-baa/doc/tree/master/zh-CN/component/pongo2.md)

### 缓存

缓存不是 baa内置依赖，作为一个常用组件实现，具体见：

[baa组件/缓存](https://github.com/go-baa/doc/tree/master/zh-CN/component/cache.md)
