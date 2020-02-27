package main

import (
	"time"

	"github.com/go-baa/baa"
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
