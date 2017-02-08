# Baa

an express Go web framework with routing, middleware, dependency injection, http context. 

Baa is ``no reflect``, ``no regexp``.

## Getting Started

Install:

```
go get -u gopkg.in/baa.v1
```

Example:

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

Run:

```
go run baa.go
```

Explore:

```
http://127.0.0.1:1323/
```

## Features

* route support static, param, group
* route support handler chain
* route support static file serve
* middleware supoort handle chain
* dependency injection support*
* context support JSON/JSONP/XML/HTML response
* centralized HTTP error handling
* centralized log handling
* whichever template engine support(emplement baa.Renderer)

## Examples

https://github.com/go-baa/example

* [blog](https://github.com/go-baa/example/tree/master/blog)

## Middlewares

* [gzip](https://github.com/baa-middleware/gzip)
* [accesslog](https://github.com/baa-middleware/accesslog)
* [recovery](https://github.com/baa-middleware/recovery)
* [session](https://github.com/baa-middleware/session)
* [static](https://github.com/baa-middleware/static)
* [requestcache](https://github.com/baa-middleware/requestcache)
* [nocache](https://github.com/baa-middleware/nocache)

## Components

* [cache](https://github.com/go-baa/cache)
* [render](https://github.com/go-baa/render)
* [pongo2](https://github.com/go-baa/pongo2)
* [router](https://github.com/go-baa/router)
* [pool](https://github.com/go-baa/pool)
* [bat](https://github.com/go-baa/bat)

