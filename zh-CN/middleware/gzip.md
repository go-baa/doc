# Baa 中间件 Gzip

`github.com/baa-middleware/gzip`

Gzip中间件提供了HTTP内容输出的Gzip配置和处理，可以减小网络传输。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/gzip"
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	if baa.Env == baa.PROD {
		app.Use(gzip.Gzip(gzip.Options{
			CompressionLevel: 9,
		}))
	}

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

在该示例中，注册了 gzip 仅在baa的产品模式开启，因为调试模式下开启gzip会导致500错误时不能正常浏览错误信息。这算不算这个BUG？

## 配置

### CompressionLevel `int`

压缩级别，参数值范围 -1 ~ 9，在 官方包中有定义，还给出了几个常量值的意义：

https://golang.org/pkg/compress/flate/#pkg-constants

```
const (
    NoCompression      = 0
    BestSpeed          = 1
    BestCompression    = 9
    DefaultCompression = -1
)
```

