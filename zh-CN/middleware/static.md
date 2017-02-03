# Baa 中间件 静态资源优先

`github.com/baa-middleware/static`

静态资源优先中间件，根据配置的`prefix`拦截路由，如果能匹配到静态资源则直接返回静态资源，否则进入路由匹配。

因为 `baa` 是一个自由的web开发框架，默认的路由行为就是一切都是 `404`，并不像 nginx/apache 等web服务器，只要有对应路径的静态文件就会显示。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/baa-middleware/static"
	"gopkg.in/baa.v1"
)

func mainStatic() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	// static
	app.Use(static.Static("/assets", "public/assets", false, nil))

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

## 配置

```
func Static(prefix, dir string, index bool, h baa.HandlerFunc) baa.HandlerFunc
```

### prefix

URI匹配前缀，如：/public, /assets

### dir

静态资源路径，可以使用绝对路径，或者相对于运行路径的路径。

### index

是否运行列出目录。

### h 

附件方法，传入一个 baa.Context 运行对输入和输出做处理。

