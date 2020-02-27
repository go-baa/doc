package main

import (
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func mainRecovery() {
	app := baa.Default()
	app.Use(recovery.Recovery())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
