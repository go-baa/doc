package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/nocache"
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func mainNocache() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())
	app.Use(nocache.New())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	}, nocache.NewFunc())

	app.Run(":1323")
}
