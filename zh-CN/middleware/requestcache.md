# Baa 中间件 请求缓存控制

`github.com/baa-middleware/requestcache`

请求缓存控制中间件可以针对全局或某个路由以请求的URI为KEY进行服务端内容缓存。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/baa-middleware/requestcache"
	"github.com/go-baa/cache"
	"gopkg.in/baa.v1"
)

func mainRequestcache() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	// request cache middleware
	app.Use(requestcache.Middleware(requestcache.Option{
		Enabled: true,
		Expires: requestcache.DefaultExpires, // 1 minute
		Headers: map[string]string{
			"X-DIY": "baa",
		},
		ContextRelated: false,
	}))

	// request cache depend cacher
	app.SetDI("cache", cache.New(cache.Options{
		Name:    "pageCache",
		Prefix:  "MyApp",
		Adapter: "memory",
		Config: map[string]interface{}{
			"bytesLimit": int64(128 * 1024 * 1024), // 128m
		},
	}))

	app.Get("/", func(c *baa.Context) {
		c.String(200, "Hello, 世界")
	})

	app.Run(":1323")
}
```

建议该中间件的注册顺序为`最后一个`，避免由于其他中间件的逻辑缓存错误的内容。

## 配置

> 本中间件依赖名为 `cache` 的 DI，需要注册一个 [Cacher](https://github.com/go-baa/cache)，cacher 的配置见：[依赖注入/缓存](https://github.com/go-baa/doc/tree/master/zh-CN/component/cache.md)

### Enabled

是否开启请求缓存控制，值类型：`bool`，默认值：`false`

### Expires

缓存过期时间，单位：秒，默认值：`600`

### Headers

附加返回头部，类型 `map[string]string`，默认值：`nil`

### ContextRelated

是否将HTTP上下文中传递的数据加入KEY组合，值类型：`bool`，默认 `false`

如果开启，通过 `c.Set()` 设置的内容将会作为缓存KEY的一部分，默认缓存的kEY是URI

## 更多姿势

### 全局使用

配置为最后一个中间件。

```
if baa.Env == baa.PROD {
	// Gzip
	b.Use(gzip.Gzip(gzip.Options{CompressionLevel: 4}))

	// Request Cache
	b.Use(requestcache.Middleware(requestcache.Option{
		Enabled: true,
		Expires: requestcache.DefaultExpires,
	}))
}
```

### 路由使用

配置为某一具体路由使用。

```
cache := requestcache.Middleware(requestcache.Option{
	Enabled: !b.Debug(),
	Expires: requestcache.DefaultExpires,
})

b.Group("/some-prefix", func() {
	// ...
}, cache)
```

### 使用多个配置

可以在不同的路由中使用不同的配置。

```
cache1 := requestcache.Middleware(requestcache.Option{
	Enabled: !b.Debug(),
	Expires: 60 * 10,
})

b.Group("/some-prefix", func() {
	// ...
}, cache1)

cache2 := requestcache.Middleware(requestcache.Option{
	Enabled: !b.Debug(),
	Expires: 60 * 30,
})

b.Group("/some-prefix-2", func() {
	// ...
}, cache2)
```
