# Baa 中间件 访问日志

访问日志中间件提供了简单的HTTP访问日志，记录了请求方法，来源IP，URL地址，响应码，发送数据，以及访问时间。

可以应付简单的日志需求，而且这个中间件只有三行代码，你可以任意定制修改。

这个中间件是学习中间件姿势的`范例`。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

建议该中间件的注册顺序为`第二`，可以更加精准的获取业务的执行时间。
