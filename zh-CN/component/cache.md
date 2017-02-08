# Baa 缓存

`github.com/go-baa/cache`

缓存组件提供了通用的缓存管理能力，采用适配器模式开发，目前支持 memory, memcache, redis 三种存储。

支持常用操作命令：Get/Set/Incr/Decr/Delete/Exist/Flush

## 使用

缓存组件是通用的，和`baa`的结合通过`DI`来引入：

```
package main

import (
    "github.com/go-baa/cache"
    "gopkg.in/baa.v1"
)

func main() {
    // new app
    app := baa.New()

    // register cache
    app.SetDI("cache", cache.New(cache.Options{
        Name:     "cache",
        Prefix:   "MyApp",
        Adapter:  "memory",
        Config:   map[string]interface{}{},
    }))

    // router
    app.Get("/", func(c *baa.Context) {
        ca := c.DI("cache").(cache.Cacher)
        ca.Set("test", "baa", 10)
        var v string
        ca.Get("test", &v)
        c.String(200, v)
    })

    // run app
    app.Run(":1323")
}
```

`memory`适配器是默认引入的，使用其他适配器需要先导入：

```
import(
    "gopkg.in/baa.v1"
    "github.com/go-baa/cache"
    _ "github.com/go-baa/cache/memcache"
    _ "github.com/go-baa/cache/redis"
)
```

### 存储 Set

### 获取 Get

### 删除 Delete

### 递增 Incr

### 递减 Decr

### 检测是否存在 Exist

### 清空缓存 Flush



## 配置

### Common

**Name**

``string``

the cache name

**Prefix**

``string``

the cache key prefix, used for isolate different cache instance/app.

**Adapter**

``string``

the cache adapter name, choose support adapter: memory, file, memcache, redis.

**Config**

``map[string]interface{}``

the cache adapter config, use a dict, values was diffrent with adapter.

### Adapter Memory

**bytesLimit**

``int64``

set the memory cache memory limit, default is 128m

**Usage**

```
app.SetDI("cache", cache.New(cache.Options{
    Name:     "cache",
    Prefix:   "MyApp",
    Adapter:  "memory",
    Config:   map[string]interface{}{
        "bytesLimit": int64(128 * 1024 * 1024), // 128m
    },
}))
```

### Adapter Memcache

**host**

``string``

memcached server host.

**port**

``string``

memcached server port.

**Usage**

```
app.SetDI("cache", cache.New(cache.Options{
    Name:     "cache",
    Prefix:   "MyApp",
    Adapter:  "memcache",
    Config:   map[string]interface{}{
        "host": "127.0.0.1",
        "port": "11211",
    },
}))
```

### Adapter Redis

**host**

``string``

redis server host.

**port**

``string``

redis server port.

**password**

``string``

redis server auth, default none.

**poolsize**

``int``

connection pool size, default 10.

**Usage**

```
app.SetDI("cache", cache.New(cache.Options{
    Name:     "cache",
    Prefix:   "MyApp",
    Adapter:  "redis",
    Config:   map[string]interface{}{
        "host":     "127.0.0.1",
        "port":     "6379",
        "password": "",
        "poolsize": 10,
    },
}))
```

