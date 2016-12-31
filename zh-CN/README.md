# Baa

一个简单高效的Go web开发框架。主要有路由、中间件，依赖注入和HTTP上下文构成。

Baa 不使用 ``反射`` 和 ``正则``，没有魔法的实现。

## 快速上手

安装:

```
go get -u gopkg.in/baa.v1
```

示例:

```
package main

import (
    "gopkg.in/baa.v1"
)

func main() {
    app := baa.New()
    app.Get("/", func(c *baa.Context) {
        c.String(200, "Hello World!")
    })
    app.Run(":1323")
}
```

## 特性

* 支持静态路由、参数路由、组路由（前缀路由/命名空间）和路由命名
* 路由支持链式操作
* 路由支持文件/目录服务
* 支持中间件和链式操作
* 支持依赖注入*
* 支持JSON/JSONP/XML/HTML格式输出
* 统一的HTTP错误处理
* 统一的日志处理
* 支持任意更换模板引擎（实现baa.Renderer接口即可）

## 示例

https://github.com/go-baa/example

* [blog](https://github.com/go-baa/example/tree/master/blog)

## 中间件

* [gzip](https://github.com/baa-middleware/gzip)
* [accesslog](https://github.com/baa-middleware/accesslog)
* [recovery](https://github.com/baa-middleware/recovery)
* [session](https://github.com/baa-middleware/session)
* [static](https://github.com/baa-middleware/static)
* [requestcache](https://github.com/baa-middleware/requestcache)
* [nocache](https://github.com/baa-middleware/nocache)

## 组件

* [cache](https://github.com/go-baa/cache)
* [render](https://github.com/go-baa/render)
* [pongo2](https://github.com/go-baa/pongo2)
* [router](https://github.com/go-baa/router)
* [pool](https://github.com/go-baa/pool)
* [bat](https://github.com/go-baa/bat)

