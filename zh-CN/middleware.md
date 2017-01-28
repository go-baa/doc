# Baa 中间件

baa 支持通过 `中间件` 机制注入请求过程，实现类似插件的功能。

根据注册的顺序依次执行，在整个执行过程中通过共享 `Context` 来实现数据传递和流程控制。

多个中间件的交替运行以 `Context.Next()` 和 `Context.Break()` 方法来控制和转换。

> 路由执行时，中间件和路由方法会组成链式操作，中间件优先路由方法先进入执行。

## 编写中间件

```
func (b *Baa) Use(m ...Middleware)
```

通过 `Use` 方法就可以注册一个中间件，接收多个 `Middleware` 类型。

```
type Middleware interface{}
```

举个例子

```
package main

import (
	"time"

	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()

	app.Use(func(c *baa.Context) {
		// 进入，记录时间
		start := time.Now()

		// 接着执行其他中间件
		c.Next()

		// 执行完其他的，最后，输出请求日志
		c.Baa().Logger().Printf("%s %s %s %v %v", c.RemoteAddr(), c.Req.Method, c.URL(false), c.Resp.Status(), time.Since(start))
	})

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

看上面的例子，整个中间件只有三句代码，要说明的是，输出日志放在了 `c.Next()` 之后执行时为了获得业务的执行时间，并没有要求必须放在那儿，只要有 `c.Next()` 就可以了。

最后，在中间件过程中，如果要中断路由操作提前退出，可以使用 `c.Break()`。

> 其实，上面的示例就是我们的 [accesslog](https://github.com/baa-middleware/accesslog) 中间件。

## 使用中间件

我们已经编写了一些常用的中间件，可以直接引入使用。

1. import
2. Use

使用示例：

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Use(accesslog.Logger())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

不同的中间件在引入的时候，可能有不同的参数，注意看各个组件部分的文档介绍。
