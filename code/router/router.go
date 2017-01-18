package main

import (
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})
	app.Post("/", func(c *baa.Context) {
		c.String(200, c.Req.Method)
	})
	app.Get("/admin", func(c *baa.Context) {
		if c.GetCookie("login_id") != "admin" {
			c.Redirect(302, "/login")
			c.Break()
		}
	}, func(c *baa.Context) {
		c.String(200, "恭喜你，看到后台了")
	})
	app.Get("/login", func(c *baa.Context) {
		c.Resp.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.SetCookie("login_id", "admin", 3600, "/")
		c.Resp.Write([]byte("登录成功，<a href=\"/admin\">点击进入后台</a>"))
	})
	app.Run(":1323")
}
