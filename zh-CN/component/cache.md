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

`func Set(key string, v interface{}, ttl int64) error`

根据 `key` 存储 `v` 的值，缓存有效期为 `ttl`秒，返回 `nil` 或者 `error`

> `ttl` 等于0，表示永不过期，如果是内存适配器，在应用重启后数据将丢失

### 获取 Get

`func Get(key string, out interface{}) error`

根据 `key` 获取存储的值赋给 `out`，返回 `nil` 或者 `error`

> 要求 `out` 为 `引用类型`，Go语言中是按值传递，如果不是引用类型外部无法获取到数据

### 删除 Delete

`func Delete(key string) error`

删除指定 `key` 的缓存数据，返回 `nil` 或者 `error`

### 递增 Incr

`func Incr(key string) (int64, error)`

根据 `key` 将存储的值 `加1` 更新存储并返回，额外返回 `nil` 或者 `error`

> `key` 的值应为 数值类型，否则将发生错误

### 递减 Decr

`func Decr(key string) (int64, error)`

根据 `key` 将存储的值 `减1` 更新存储并返回，额外返回 `nil` 或者 `error`

> `key` 的值应为 数值类型，否则将发生错误

### 检测是否存在 Exist

`func Exist(key string) bool`

根据 `key` 检查是否存在有效的缓存数据，返回 `true` 表示存在，`false` 表示不存在

### 清空缓存 Flush

`func Flush() error`

清空缓存中的数据

## 配置

### Name `string`

缓存实例名称

### Prefix `string`

缓存索引前缀

### Adapter `string`

适配器名称，目前支持 memory, memcache, redis 三种存储

### Config `map[string]interface{}`

适配器配置，不同的适配器有不同的配置。

### 适配器 Memory

#### bytesLimit `int64`

内存适配器，只有一个配置参数，内存大小限制，单位 字节，默认为 `128m`

#### 使用示例

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

### 适配器 Memcache

#### host `string`

memcached 服务器IP地址，默认为 `127.0.0.1`

#### port `string`

memcached 服务器端口，默认为 `11211`

#### 使用示例

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

### 适配器 Redis

#### host `string`

redis 服务器IP地址，默认为 `127.0.0.1`

#### port `string`

redis 服务器端口，默认为 `6379`

#### password `string`

redis 服务器连接密码，默认 `空`

#### poolsize `int`

redis 库连接池限制，默认保持 `10` 个连接

#### 使用示例

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
