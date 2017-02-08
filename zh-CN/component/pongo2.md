# Baa pongo2

`https://github.com/go-baa/pongo2`

pongo2 是 一个将 [Pongo2](https://github.com/flosch/pongo2) 模板引擎应用到 baa的助手库。

## 使用

通过 `DI` 注册 `render` 即可。

```
package main

import (
    "github.com/go-baa/pongo2"
    "gopkg.in/baa.v1"
)

func main() {
    // new app
    app := baa.New()

    // register pongo2 render
    // render is template DI for baa, must be this name.
    app.SetDI("render", pongo2.New(pongo2.Options{
        Baa:        b,
        Root:       "templates/",
        Extensions: []string{".html"},
        Functions:  map[string]interface{}{},
        Context: map[string]interface{}{
            "SITE_NAME": "Yet another website",
        },
    }))

    // router
    app.Get("/", func(c *baa.Context) {
        c.HTML(200, "index")
    })

    // run app
    app.Run(":1323")
}
```

## 配置 `pongo2.Options`

### Baa `*baa.Baa`

render 需要传递 baa 实例

### Root `string`

模板目录，指定模板目录，`c.HTML` 渲染模板时将从该目录下查找目录，不需要在渲染是指定目录名。

### Extensions `[]string`

模板文件扩展名，可以指定多个，渲染模板时无需指定扩展名，将根据配置寻找对应的文件。

### Functions `map[string]interface{}`

扩展函数，参考：[FuncMap](https://godoc.org/html/template#FuncMap)

### Context `map[string]interface{}`

预置数据，模板变量

## 扩展语法

### 输出变量

```
{{ name }}
```

### 加载模板片段

```
{% include "path/to/tpl.html" %}
```

带参数加载

```
{% include "relative/path/to/tpl.html" with foo=var %}
{% include "relative/path/to/tpl.html" with foo="bar" %}
```

> 嵌入的模板接收的参数将作为 `string` 类型。

### 条件语句

```
{% if vara %}
{% elif varb %}
{% else %}
{% endif %}
```

### 循环语句

```
{% for item in items %}
{{ forloop.Counter }} {{ forloop.Counter0 }} {{ forloop.First }} {{ forloop.Last }} {{ forloop.Revcounter }} {{ forloop.Revcounter0 }}
{{ item }}
{% endfor %}
```

### 内置过滤器

* escape
* safe
* escapejs
* add
* addslashes
* capfirst
* center
* cut
* date
* default
* default_if_none
* divisibleby
* first
* floatformat
* get_digit
* iriencode
* join
* last
* length
* length_is
* linebreaks
* linebreaksbr
* linenumbers
* ljust
* lower
* make_list
* phone2numeric
* pluralize
* random
* removetags
* rjust
* slice
* stringformat
* striptags
* time
* title
* truncatechars
* truncatechars_html
* truncatewords
* truncatewords_html
* upper
* urlencode
* urlize
* urlizetrunc
* wordcount
* wordwrap
* yesno
* float
* integer

### extends / block / macro and so on ...

更多内容，参见 [django](https://docs.djangoproject.com/en/dev/ref/templates/language/).
