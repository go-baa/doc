# Baa render

`https://github.com/go-baa/render`

render 是一个 增强型模板引擎，相对 baa 内置的模板引擎有如下特性：

- 模板容器，可以使用 `template` 语法加载模板片段.
- 模板缓存，在程序启动时即加载所有模板到内存中执行预编译，并且可以感知文件变化自动重新编译.
- 可定制模板目录，模板文件扩张名，支持扩充模板函数.

## 使用

通过 `DI` 注册 `render` 即可。

```
package main

import (
    "github.com/go-baa/render"
    "gopkg.in/baa.v1"
)

func main() {
    // new app
    app := baa.New()

    // register render
    // render is template DI for baa, must this name.
    app.SetDI("render", render.New(render.Options{
        Baa:        app,
        Root:       "templates/",
        Extensions: []string{".html", ".tmpl"},
    }))

    // router
    app.Get("/", func(c *baa.Context) {
        c.HTML(200, "index")
    })

    // run app
    app.Run(":1323")
}
```

## 配置 `render.Options` 

### Baa `*baa.Baa`

render 需要传递 baa 实例

### Root `string`

模板目录，指定模板目录，`c.HTML` 渲染模板时将从该目录下查找目录，不需要在渲染是指定目录名。

### Extensions `[]string`

模板文件扩展名，可以指定多个，渲染模板时无需指定扩展名，将根据配置寻找对应的文件。

### Functions `map[string]interface{}`

扩展函数，参考：[FuncMap](https://godoc.org/html/template#FuncMap)

## 扩展语法

### 加载模板片段 

```
{{ template "share/footer" }}
```
