package main

import (
	"github.com/go-baa/baa"
	"github.com/go-baa/cache"
)

func main() {
	// new app
	app := baa.New()

	// register cache
	app.SetDI("cache", cache.New(cache.Options{
		Name:    "cache",
		Prefix:  "MyApp",
		Adapter: "memory",
		Config:  map[string]interface{}{},
	}))

	// router
	app.Get("/", func(c *baa.Context) {
		ca := c.DI("cache").(cache.Cacher)
		ca.Set("test", "baa", 10)
		var v string
		ca.Get("test", &v)
		c.String(200, v)
	})

	// run app
	app.Run(":1323")
}
