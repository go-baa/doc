# Baa router

`https://github.com/go-baa/router`

路由项目提供了额外的路由器实现。

* regexp 完全使用正则表达式的路由器，在复杂路由多为正则表达式的情况下，效率稍好。
* regtree 使用正则表达式和基数树组合优化的路由器，在多数路由为普通路由，少量正则路由条目的情况下，效率较好。

如果没有大量的正则表达式路由条目，建议使用 `regtree` 作为默认的 `正则路由`。

## 使用

通过 `DI` 替换掉全局的 `router` 即可。

```
package main

import (
    "github.com/go-baa/router/regtree"
    "github.com/go-baa/baa"
)

func main() {
    app := baa.Default()
    app.SetDI("router", regtree.New(app))

    app.Get("/view-:id(\\d+)", func(c *baa.Context) {
        c.String(200, c.Param("id"))
    })
    app.Get("/view-:id(\\d+)/project", func(c *baa.Context) {
        c.String(200, c.Param("id")+"/project")
    })
    app.Run(":1323")
}
```
