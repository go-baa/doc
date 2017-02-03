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
    * [日志](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#日志)
    * [路由](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#路由)
    * [模板](https://github.com/go-baa/doc/tree/master/zh-CN/di.md#模板)
    * [缓存](https://github.com/go-baa/doc/tree/master/zh-CN/component/cache.md)
* [HTTP上下文](https://github.com/go-baa/doc/tree/master/zh-CN/context.md)
    * [Context](*)
    * [路由参数](*)
    * [请求参数](*)
    * [Cookie](*)
    * [内容输出](*)
    * [模板语法](*)
        * [pongo2](*)
* [日志](*)
    * [日志接口](*)
    * [日志方法](*)
* [工具](*)
    * [bat](*)
    * [pool](*)
    * [gRPC](*)
* [数据库](*)
    * [gorm](*)
    * [xorm](*)
    * [mgo](*)
* [工程化](*)
    * [目录结构](*)
    * [配置文件](*)
    * [初始化](*)
    * [控制器](*)
    * [路由](*)
    * [模型](*)

## 快速上手

安装:

```
go get -u gopkg.in/baa.v1
```

代码:

```
// baa.go

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

运行：

```
go run baa.go
```

## 示例

https://github.com/go-baa/example

* [博客](https://github.com/go-baa/example/tree/master/blog)
* [websocket](http://github.com/go-baa/example/tree/master/websocket)
