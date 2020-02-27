package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func mainAccesslog() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
