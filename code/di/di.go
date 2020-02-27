package main

import (
	"log"
	"os"

	"github.com/go-baa/baa"
)

func main() {
	app := baa.Default()
	app.SetDI("logger", log.New(os.Stderr, "[BaaDI] ", log.LstdFlags))

	app.Get("/", func(c *baa.Context) {
		// use di
		logger := c.DI("logger").(*log.Logger)
		logger.Println("i am use logger di")

		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
