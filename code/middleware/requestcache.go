package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/baa-middleware/requestcache"
	"github.com/go-baa/baa"
	"github.com/go-baa/cache"
)

func mainRequestcache() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	// request cache middleware
	app.Use(requestcache.Middleware(requestcache.Option{
		Enabled: true,
		Expires: requestcache.DefaultExpires, // 1 minute
		Headers: map[string]string{
			"X-DIY": "baa",
		},
		ContextRelated: false,
	}))

	// request cache depend cacher
	app.SetDI("cache", cache.New(cache.Options{
		Name:    "pageCache",
		Prefix:  "MyApp",
		Adapter: "memory",
		Config: map[string]interface{}{
			"bytesLimit": int64(128 * 1024 * 1024), // 128m
		},
	}))

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
