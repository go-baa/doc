package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/gzip"
	"github.com/baa-middleware/recovery"
	"github.com/go-baa/baa"
)

func mainGzip() {
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
