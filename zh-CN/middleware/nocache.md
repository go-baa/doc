# Baa 中间件 强制不缓存

`github.com/baa-middleware/nocache`

强制不缓存中间件给所有请求增加了NO-CACHE Header头，防止浏览器或CDN缓存。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/nocache"
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())
	app.Use(nocache.New())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	}, nocache.NewFunc())

	app.Run(":1323")
}

```

强制不缓存中间件提供了两个方法：

```
nocache.New()
```

返回一个中间件，用于中间件注册。

```
nocache.NewFunc()
```

返回一个路由处理函数，用于给指定的路由请求增加NO-CACHE设定。
