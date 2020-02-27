package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/baa-middleware/session"
	"github.com/go-baa/baa"
)

func mainSession() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	// session config
	redisOptions := session.RedisOptions{}
	redisOptions.Addr = "127.0.0.1:6379"
	redisOptions.Prefix = "Prefix:"

	memoryOptions := session.MemoryOptions{
		BytesLimit: 1024 * 1024,
	}

	app.Use(session.Middleware(session.Options{
		Name: "BAASESSIONID",
		Provider: &session.ProviderOptions{
			Adapter: "memory", // redis / memory
			Config:  memoryOptions,
		},
	}))

	app.Get("/", func(c *baa.Context) {
		// get the session handler
		session := c.Get("session").(*session.Session)
		session.Set("hi", "baa")

		c.String(200, "Hello, 世界")
	})

	app.Get("/session", func(c *baa.Context) {
		// get the session handler
		session := c.Get("session").(*session.Session)
		c.String(200, "SessionID: "+session.ID()+", hi, "+session.Get("hi").(string))
	})

	app.Run(":1323")
}
