# Baa 中间件 错误恢复

错误恢复中间件提供了当业务逻辑`Panic`时记录日志和返回500错误，并且`recovey`Go的程序，防止一个业务中的错误导致应用崩溃。

## 使用

```
package main

import (
	"github.com/baa-middleware/recovery"
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

建议该中间件的注册顺序为`第一`，防止其他中间件本身就有错误导致应用崩溃。
