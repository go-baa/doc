# Baa 路由

baa 基于 http resetfull 模式设计了路由管理器，提供了常规路由，参数路由，文件路由，静态目录路由，还有组路由。

## 常规路由

```
func (b *Baa) Get(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Head(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Post(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Put(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Delete(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Patch(pattern string, h ...HandlerFunc) RouteNode
func (b *Baa) Options(pattern string, h ...HandlerFunc) RouteNode
```

* pattern URI路径
* HandlerFunc 执行方法，允许多个，按照设定顺序进行链式处理。

除了以上几个标准方法，还支持多个method设定的路由姿势：

```
func (b *Baa) Route(pattern, methods string, h ...HandlerFunc) RouteNode
func (b *Baa) Any(pattern string, h ...HandlerFunc) RouteNode
```

## 路由语法

静态路由
参数路由
正则路由

## 组路由

func (b *Baa) Group(pattern string, f func(), h ...HandlerFunc)



### 链式处理

一个URL请求可以先处理A，根据A的结果再执行B。

> 举个例子：
>
> 一个URL要先判断你登录过才可以访问，就可以设定两个Handler，第一个 判断是否登录，如果没登录就调到登录界面，否则继续执行第二个真正的内容。

使用示例：

```
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
```

## 命名路由

func (n *Node) Name(name string)
func (b *Baa) URLFor(name string, args ...interface{}) string


## 文件路由

func (b *Baa) Static(prefix string, dir string, index bool, h HandlerFunc)
func (b *Baa) StaticFile(pattern string, path string) RouteNode

## 错误路由

func (b *Baa) SetError(h ErrorHandleFunc)
func (b *Baa) SetNotFound(h HandlerFunc)

## Websocket

func (b *Baa) Websocket(pattern string, h func(*websocket.Conn)) RouteNode

