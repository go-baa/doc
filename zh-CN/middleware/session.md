# Baa 中间件 session

`github.com/baa-middleware/session`

Go中的HTTP包默认提供了对cookie的管理，但是session并没有，需要各显神通，这就是一个简单的实现。

session 中间件提供了对session的管理和使用，支持 redis/memory 适配器。

## 使用

```
package main

import (
	"github.com/baa-middleware/accesslog"
	"github.com/baa-middleware/recovery"
	"github.com/baa-middleware/session"
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.Default()
	app.Use(recovery.Recovery())
	app.Use(accesslog.Logger())

	// session config
	redisOptions := session.RedisOptions{}
	redisOptions.Addr = "127.0.0.1:6379"
	redisOptions.Prefix = "Prefix:"

	memoryOptions := session.MemoryOptions{
		BytesLimit: 1024 * 1024,
	}

	app.Use(session.Middleware(session.Options{
		Name: "BAASESSIONID",
		Provider: &session.ProviderOptions{
			Adapter: "memory", // redis / memory
			Config:  memoryOptions,
		},
	}))

	app.Get("/", func(c *baa.Context) {
		// get the session handler
		session := c.Get("session").(*session.Session)
		session.Set("hi", "baa")

		c.String(200, "Hello, 世界")
	})

	app.Get("/session", func(c *baa.Context) {
		// get the session handler
		session := c.Get("session").(*session.Session)
		c.String(200, "SessionID: "+session.ID()+", hi, "+session.Get("hi").(string))
	})

	app.Run(":1323")
}
```

在该示例中，注册了 gzip 仅在baa的产品模式开启，因为调试模式下开启gzip会导致500错误时不能正常浏览错误信息。这算不算这个BUG？

## 配置

### Name `string`

Session 的名称

### IDLength `int`

SessionID 的长度，默认 `16` 位

### Provider `*ProviderOptions`

Session的存储器选项

#### Adapter `string`

存储器的适配器类型，支持 `redis` 和 `memory`

> 注意：`memory` 适配器类型，如果应用重启session将全部失效，在集群模式下也不能共享session。

#### Config `interface{}`

存储器的适配器配置，不同的适配器类型应使用不同的配置, `session.RedisOptions{}` 或者 `session.MemoryOptions{}`

### Cookie `*CookieOptions`

Session的Cookie配置

#### Domain `string`

Cookie有效域名，默认 空

#### Path `string`

Cookie有效路径，默认 `/`

#### Secure `bool`

Cookie是否加密，默认 `false`

#### LifeTime `int64`

Cookie有效期，默认 `0`，即是一个浏览器会话类型，关闭响应的页面即失效。

#### HttpOnly `bool`

Cookie是否仅浏览器有效，默认 `false`

### GCInterval `int64`

Session GC频率，仅在 `memory` 存储器类型下有效

### MaxLifeTime `int64`

Session有效期，默认 `0`
