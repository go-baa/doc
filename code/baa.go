package main

import (
	"github.com/go-baa/baa"
)

func main() {
	app := baa.Default()
	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})
	app.Run(":1323")
}
