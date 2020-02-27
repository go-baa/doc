package main

import (
	"fmt"

	"github.com/go-baa/baa"
	"github.com/go-baa/router/regtree"
)

func main() {
	app := baa.New()
	app.SetDI("router", regtree.New(app))

	// normal
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

	// static route
	app.Get("/foo", func(c *baa.Context) {
		c.String(200, c.URL(true))
	})
	app.Get("/bar", func(c *baa.Context) {
		c.String(200, c.URL(true))
	})

	// param route
	app.Get("/user/:id", func(c *baa.Context) {
		c.String(200, "My user id is: "+c.Param("id"))
	})
	app.Get("/user/:id/project/:pid", func(c *baa.Context) {
		id := c.ParamInt("id")
		pid := c.ParamInt("pid")
		c.String(200, fmt.Sprintf("user id: %d, project id: %d", id, pid))
	})

	// regexp route
	app.Get("/user-:id(\\d+)", func(c *baa.Context) {
		c.String(200, "My user id is: "+c.Param("id"))
	})
	app.Get("/user-:id(\\d+)-project-:pid(\\d+)", func(c *baa.Context) {
		id := c.ParamInt("id")
		pid := c.ParamInt("pid")
		c.String(200, fmt.Sprintf("user id: %d, project id: %d", id, pid))
	})

	// group
	app.Group("/group", func() {
		app.Get("/", func(c *baa.Context) {
			c.String(200, "我是组的首页")
		})
		app.Group("/user", func() {
			app.Get("/", func(c *baa.Context) {
				c.String(200, "我是组的用户")
			})
			app.Get("/:id", func(c *baa.Context) {
				c.String(200, "in group, user id: "+c.Param("id"))
			})
		})
		app.Get("/:gid", func(c *baa.Context) {
			c.String(200, "in group, group id: "+c.Param("gid"))
		})
	}, func(c *baa.Context) {
		// 我是组内的前置检测，过不了我这关休想访问组内的资源
	})

	app.Run(":1323")
}
