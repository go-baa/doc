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

session name

### IDLength `int`

session id length, default is `16`

### Provider `*ProviderOptions`

provider options

#### Adapter `string`

provider adapter name, currently support `redis` and `memory`

#### Config `interface{}`
 
provider adapter config, each adapter has its own config

### Cookie `*CookieOptions`

session cookie options

#### Domain `string`

cookie domain, default is `''`

#### Path `string`

cookie path, default is `/`

#### Secure `bool`

#### LifeTime `int64`

cookie life time, default is `0`, known as session cookie

#### HttpOnly `bool`

### GCInterval `int64`

garbage collection run interval, used for `memory` adapter only

### MaxLifeTime `int64`

After this number of seconds, stored data will be seen as 'garbage' and cleaned up by the garbage collection process


