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
