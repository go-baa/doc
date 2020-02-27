package main

import baa "github.com/go-baa/baa"

func main() {
	app := baa.New()
	app.Get("/", func(c *baa.Context) {
		c.Set("title", "this is title")
		c.Set("content", "this is content")
		c.Set("show", true)
		c.Set("list", []string{"111", "222", "333"})
		c.HTML(200, "template/index.html")
	})
	app.Run(":1323")
}
