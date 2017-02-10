# Baa

一个简单高效的Go web开发框架。主要有路由、中间件，依赖注入和HTTP上下文构成。

Baa 不使用 ``反射`` 和 ``正则``，没有魔法的实现。

## 特性

* 支持静态路由、参数路由、组路由（前缀路由/命名空间）和路由命名
* 路由支持链式操作
* 路由支持文件/目录服务
* 中间件支持链式操作
* 支持依赖注入*
* 支持JSON/JSONP/XML/HTML格式输出
* 统一的HTTP错误处理
* 统一的日志处理
* 支持任意更换模板引擎（实现baa.Renderer接口即可）

## 文档目录

* [Baa核心](https://github.com/go-baa/doc/tree/master/zh-CN/baa.md)
* [路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md)
    * [常规路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#常规路由)
    * [路由语法](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#路由语法)
        * [静态路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#静态路由)
        * [参数路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#参数路由)
        * [正则路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#正则路由)
    * [路由选项](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#路由选项)
    * [组路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#组路由)
    * [链式处理](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#链式处理)
    * [命名路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#命名路由)
    * [文件路由](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#文件路由)
    * [自定义错误](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#自定义错误)
        * [500错误](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#500错误)
        * [404错误](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#404错误)
    * [Websocket](https://github.com/go-baa/doc/tree/master/zh-CN/router.md#websocket)
* [中间件](https://github.com/go-baa/doc/tree/master/zh-CN/middleware.md)
    * [编写中间件](https://github.com/go-baa/doc/blob/master/zh-CN/middleware.md#编写中间件)
    * [使用中间件](https://github.com/go-baa/doc/blob/master/zh-CN/middleware.md#使用中间件)
        * [错误恢复](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/recovery.md)
        * [访问日志](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/accesslog.md)
        * [gzip](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/gzip.md)
        * [session](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/session.md)
        * [静态资源优先](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/static.md)
        * [请求缓存控制](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/requestcache.md)
        * [强制不缓存](https://github.com/go-baa/doc/blob/master/zh-CN/middleware/nocache.md)
* [依赖注入 DI](https://github.com/go-baa/doc/tree/master/zh-CN/di.md)
    * [注册](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#注册)
    * [使用](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#使用)
        * [日志](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#日志)
        * [路由](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#路由)
        * [模板](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#模板)
        * [缓存](https://github.com/go-baa/doc/tree/master/zh-CN/component/cache.md)
* [HTTP上下文](https://github.com/go-baa/doc/tree/master/zh-CN/context.md)
    * [Request](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#request)
        * [URL参数](https://github.com/go-baa/doc/blob/master/zh-CN/context.md#url参数)
        * [路由参数](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#路由参数)
        * [Cookie](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#cookie)
        * [文件上传](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#文件上传)
    * [Response](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#response)
        * [数据存储](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#数据存储)
        * [内容输出](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#内容输出)
        * [有用的函数](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#有用的函数)
    * [模板渲染](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板渲染)
        * [模板语法](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板语法)
        * [模板接口](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板接口)
        * [render](https://github.com/go-baa/doc/tree/master/zh-CN/component/render.md)
        * [pongo2](https://github.com/go-baa/doc/tree/master/zh-CN/component/pongo2.md)
* [日志](https://github.com/go-baa/doc/tree/master/zh-CN/log.md)
    * [日志接口](https://github.com/go-baa/doc/tree/master/zh-CN/log.md#日志接口)
    * [日志方法](https://github.com/go-baa/doc/tree/master/zh-CN/log.md#日志方法)
* [数据库](https://github.com/go-baa/doc/tree/master/zh-CN/database.md)
    * [gorm](http://jinzhu.me/gorm/)
    * [xorm](http://xorm.io/)
    * [mgo](https://labix.org/mgo)
* [组件](https://github.com/go-baa/doc/tree/master/zh-CN/component.md)
    * [bat](https://github.com/go-baa/doc/tree/master/zh-CN/component/bat.md)
    * [cache](https://github.com/go-baa/doc/tree/master/zh-CN/component/cache.md)
    * [pongo2](https://github.com/go-baa/doc/tree/master/zh-CN/component/pongo2.md)
    * [pool](https://github.com/go-baa/doc/tree/master/zh-CN/component/pool.md)
    * [render](https://github.com/go-baa/doc/tree/master/zh-CN/component/render.md)
    * [router](https://github.com/go-baa/doc/tree/master/zh-CN/component/router.md)
    * [log](https://github.com/go-baa/doc/tree/master/zh-CN/component/log.md)
    * [setting](https://github.com/go-baa/doc/tree/master/zh-CN/component/setting.md)
* [工程化](https://github.com/go-baa/doc/tree/master/zh-CN/project.md)
    * [目录结构](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#目录结构)
    * [控制器](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#控制器)
    * [数据模型](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#数据模型)
    * [配置文件](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#配置文件)
    * [模板](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#模板)
    * [静态资源](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#静态资源)
    * [打包发布](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#打包发布)
    * [依赖管理](ttps://github.com/go-baa/doc/tree/master/zh-CN/project.md#依赖管理)


## 快速上手

### 安装

```
go get -u gopkg.in/baa.v1
```

### 代码

```
// baa.go
package main

import (
    "gopkg.in/baa.v1"
)

func main() {
    app := baa.New()
    app.Get("/", func(c *baa.Context) {
        c.String(200, "Hello, 世界")
    })
    app.Run(":1323")
}
```

### 运行

```
go run baa.go
```

### 浏览

```
http://127.0.0.1:1323/
```

### 使用中间件

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

## 示例

https://github.com/go-baa/example

* [blog](https://github.com/go-baa/example/tree/master/blog)
* [api](http://github.com/go-baa/example/tree/master/api)
* [websocket](http://github.com/go-baa/example/tree/master/websocket)
