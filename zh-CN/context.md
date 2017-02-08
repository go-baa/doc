# Baa Context 

Context是HTTP上下文的意思，封装了`输入`、`输出`，提供请求处理，结果响应，模板渲染等相关操作。

## Request

`c.Req`

Request 中包含了所有的请求数据，是标准包中 `http.Request` 的封装，可以通过 `c.Req` 访问原生的请求结构。

### URL参数

`func (c *Context) Posts() map[string]interface{}`

获取所有的POST 和 GET 数据，返回一个字典，字典的索引为表单字段的名称，值为字段的值，值如果只有一个则为 `string` 类型，值有多个则为 `[]string` 类型。

`func (c *Context) Querys() map[string]interface{}`

获取所有的 GET 数据，返回一个字典，字典的索引为表单字段的名称，值为字段的值，值如果只有一个则为 `string` 类型，值有多个则为 `[]string` 类型。

`func (c *Context) Query(name string) string`

根据 `name` 获取一个字段的值，返回 `string` 类型，包含 POST 和 GET 数据。

`func (c *Context) QueryStrings(name string) []string`

根据 `name` 获取一个字段的多个值，返回 `[]string` 类型。

`func (c *Context) QueryEscape(name string) string`

根据 `name` 获取一个字段的值，escapre编码后返回 `string` 类型。

`func (c *Context) QueryTrim(name string) string`

根据 `name` 获取一个字段的值，去除两端空白后返回 `string` 类型。

`func (c *Context) QueryBool(name string) bool`

根据 `name` 获取一个字段的值，并强制转化为 `bool` 类型 返回。

`func (c *Context) QueryFloat(name string) float64`

根据 `name` 获取一个字段的值，并强制转化为 `float64` 类型 返回。

`func (c *Context) QueryInt(name string) int`

根据 `name` 获取一个字段的值，并强制转化为 `int` 类型 返回。

`func (c *Context) QueryInt32(name string) int32`

根据 `name` 获取一个字段的值，并强制转化为 `int32` 类型 返回。

`func (c *Context) QueryInt64(name string) int64`

根据 `name` 获取一个字段的值，并强制转化为 `int64` 类型 返回。

### 路由参数

`func (c *Context) Param(name string) string`

根据 `name` 获取一个路由参数的值，返回 `string` 类型。

`func (c *Context) Params() map[string]string`

返回所有的路由参数组成的字典，字典的索引是参数名称。

`func (c *Context) ParamBool(name string) bool`

根据 `name` 获取一个路由参数的值，并强制转化为 `bool` 类型 返回。

`func (c *Context) ParamFloat(name string) float64`

根据 `name` 获取一个路由参数的值，并强制转化为 `float64` 类型 返回。

`func (c *Context) ParamInt(name string) int`

根据 `name` 获取一个路由参数的值，并强制转化为 `int` 类型 返回。

`func (c *Context) ParamInt32(name string) int32`

根据 `name` 获取一个路由参数的值，并强制转化为 `int32` 类型 返回。

`func (c *Context) ParamInt64(name string) int64`

根据 `name` 获取一个路由参数的值，并强制转化为 `int64` 类型 返回。

### Cookie

`func (c *Context) GetCookie(name string) string`

根据 `name` 获取一个Cookie的值，返回 `string` 类型。

`func (c *Context) GetCookieBool(name string) bool`

根据 `name` 获取一个Cookie的值，并强制转化为 `bool` 类型 返回。

`func (c *Context) GetCookieFloat64(name string) float64`

根据 `name` 获取一个Cookie的值，并强制转化为 `float64` 类型 返回。

`func (c *Context) GetCookieInt(name string) int`

根据 `name` 获取一个Cookie的值，并强制转化为 `int` 类型 返回。

`func (c *Context) GetCookieInt32(name string) int32`

根据 `name` 获取一个Cookie的值，并强制转化为 `int32` 类型 返回。

`func (c *Context) GetCookieInt64(name string) int64`

根据 `name` 获取一个Cookie的值，并强制转化为 `int64` 类型 返回。

`func (c *Context) SetCookie(name string, value string, others ...interface{})`

设置 名称为 `name` 值为 `value` 的Cookie。

`setCookie` 还可以指定更多Cookie参数，通过可变长度的 `others` 参数可以依次指定：

```
SetCookie(<name>, <value>, <max age>, <path>, <domain>, <secure>, <http only>)
```

使用姿势：

```
c.SetCookie("mykey", "myvalue")
c.SetCookie("mykey", "myvalue", 3600)
c.SetCookie("mykey", "myvalue", 3600, "/")
c.SetCookie("mykey", "myvalue", 3600, "/", ".vodjk.com")
c.SetCookie("mykey", "myvalue", 3600, "/", ".vodjk.com", true)
c.SetCookie("mykey", "myvalue", 3600, "/", ".vodjk.com", true, true)
```

> 可变参数需依次指定，不能跳过中间的参数

### 文件上传

`func (c *Context) GetFile(name string) (multipart.File, *multipart.FileHeader, error)`

通过 `multipart/form-data` 表单上传的文件，可以通过 `c.GetFile` 指定文件字段获得到文件结构 
[multipart.File](https://godoc.org/mime/multipart#File) 和 [multipart.FileHeader](https://godoc.org/mime/multipart#FileHeader)，

然后可以利用这些结构进行文件大小检测，文件类型检测，文件存储等。

举个栗子：

```
func Upload(c *baa.Context) {
	file, header, err := c.GetFile("file1")
	if err != nil {
		c.Error(errors.New("没有文件被上传"))
		return
	}
	defer file.Close()

	savedTo := "savedFile.jpg"
	newFile, err := os.Create(savedTo)
	if err != nil {
		c.Error(errors.New("文件创建失败"))
		return
	}
	defer newFile.Close()

	size, err := io.Copy(newFile, file)
	msg := fmt.Sprintf("fileName: %s, savedTo: %s, size: %d, err: %v", header.Filename, savedTo, size, err)
	fmt.Println(msg)

	c.String(200, msg)
}

```

`func (c *Context) SaveToFile(name, savePath string) error`

快速保存指定的文件字段 `name` 到指定的路径 `savePath` 并返回 `nil` 或者 `error`

还是上面的例子，改一下：

```
func Upload2(c *baa.Context) {
	savedTo := "savedFile2.jpg"
	err := c.SaveToFile("file1", savedTo)
	if err != nil {
		c.Error(err)
		return
	}
	c.String(200, savedTo)
}
```

## Response

`c.Resp`

Response 中用于处理结果输出，是标准包中 `http.ResponseWriter` 的封装，可以通过 `c.Resp` 访问原生的接口。

### 数据存储

`Context` 提供了临时存储可用于整个请求的生命周期中。

`func (c *Context) Set(key string, v interface{})`

根据 `key` 设置一个值 `v` 到 `Context` 存储中。

`func (c *Context) Get(key string) interface{}`

根据 `key` 从`Context`存储中获取一个值并返回。

`func (c *Context) Gets() map[string]interface{}`

获取`Context`中存储的所有值，返回由这些值组成的字典。

举个例子：

```
package main 

import (
    "gopkg.in/baa.v1"
)

func main() {
    app := baa.New()
    app.Get("/", func(c *baa.Context) {
        c.Set("mykey", "myvalue")
    }, func(c *baa.Context) {
        val := c.Get("mykey")
        fmt.Println(val) // myvalue

        c.String(200, val.(string))
    })

    app.Run(":1323")
}
```

### 内容输出

baa 提供了多种形式的内容输出。

`func (c *Context) String(code int, s string)`

设定输出的 http code 为 `code`，并输出一个字符串 `s`。

`func (c *Context) Text(code int, s []byte)`

设定输出的 http code 为 `code`，并输出一个字节切片 `s`。

`func (c *Context) JSON(code int, v interface{})`

设定输出的 http code 为 `code`，设定内容类型为 `application/json`， 把 结构 `v` 使用JSON编码后输出。

`func (c *Context) JSONP(code int, callback string, v interface{})`

设定输出的 http code 为 `code`，设定内容类型为 `application/json`， 把 结构 `v` 使用JSON编码，并结合 `callback`参数输出。

`func (c *Context) JSONString(v interface{}) (string, error)`

把 结构 `v` 使用JSON编码后返回。

`func (c *Context) XML(code int, v interface{})`

设定输出的 http code 为 `code`，设定内容类型为 `application/json`， 把 结构 `v` 使用XML编码后输出。

## 有用的函数

`func (c *Context) Baa() *Baa`

`func (c *Context) Body() *RequestBody`
`func (c *Context) Break()`
`func (c *Context) Error(err error)`

`func (c *Context) Next()`
`func (c *Context) NotFound()`
`func (c *Context) IsAJAX() bool`
`func (c *Context) IsMobile() bool`

`func (c *Context) Redirect(code int, url string) error`
`func (c *Context) Referer() string`
`func (c *Context) RemoteAddr() string`
`func (c *Context) URL(hasQuery bool) string`
`func (c *Context) UserAgent() string`

## 模板渲染

baa 集成一个简单的模板渲染，使用Go标准库的 [template语法](https://godoc.org/html/template)。

baa 的模板渲染使用 `Context`存储 中的数据作为模板变量。

`func (c *Context) Fetch(tpl string) ([]byte, error)`

根据 `tpl` 模板文件路径，使用 `Context`存储中的数据，渲染模板并返回渲染后的内容。

`func (c *Context) Render(code int, tpl string)`

设定输出的 http code 为 `code`，设定内容类型为 `text/html`， 渲染模板 `tpl` 并直接输出。

> 内部就是调用的 `c.Fetch` 然后 把内容输出

`func (c *Context) HTML(code int, tpl string)`

`c.Render`的一个别名，用起来更顺手。

举个例子：

```
package main

import baa "gopkg.in/baa.v1"

func main() {
	app := baa.New()
	app.Get("/", func(c *baa.Context) {
		c.Set("title", "this is title")
		c.Set("content", "this is content")
		c.HTML(200, "template/index.html")
	})
	app.Run(":1323")
}

```

```
<!-- template/index.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
</head>
<body>

    {{ .content}}
    
</body>
</html>
```

### 模板语法

以下仅做简单介绍，完整文档请见官方 [html/template](https://godoc.org/html/template)。

* 模板语法以一对 `双大括号` 包裹
* 变量都是以 `.` 开始，`.` 代表所有数据组成的结构
* 结构体中的字段用 `.` 表示子集

#### 输出变量

```
{{ .var }}
{{ .user.id }}
```

#### 条件语句

```
{{ if .show }}
    i want show!
{{ else }}
    i was hidden
{{ end }}
```

#### 循环语句

```
{{ range .list }}
    {{ .id }}
    {{ .name }}
{{ end }}
```

如果循环体不是结构体，比如就是一个字符串数组，直接用 `.` 即可输出：

```
{{ range .strs }}
    {{ . }}
{{ end }}
```

### 模板接口

baa 中运行通过 `DI` 替换模板引擎，只要实现 `baa.Renderer` 接口即可。

Renderer

```
type Renderer interface {
    Render(w io.Writer, tpl string, data interface{}) error
}
```

渲染接口只有一个方法 `Render`，该方法 接收三个参数：

#### w `io.Writer` 
 
一个可写入的类型，用于写入渲染后的数据，这里其实就是 `c.Resp` 。
 
#### tpl `string`

模板文件路径

#### data `interface{}`

向模板传递的数据（模板变量），这里其实传递过来的就是 `c.Gets` 的结果，是一个 `map[string]interface{}` 类型。

### render

`https://github.com/go-baa/render`

render 是一个 增强型模板引擎，相对 baa 内置的模板引擎有如下特性：

- 模板容器，可以使用 `template` 语法加载模板片段.
- 模板缓存，在程序启动时即加载所有模板到内存中执行预编译，并且可以感知文件变化自动重新编译.
- 可定制模板目录，模板文件扩张名，支持扩充模板函数.

#### 使用

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

#### 配置 `render.Options` 

##### Baa `*baa.Baa`

render 需要传递 baa 实例

##### Root `string`

模板目录，指定模板目录，`c.HTML` 渲染模板时将从该目录下查找目录，不需要在渲染是指定目录名。

##### Extensions `[]string`

模板文件扩展名，可以指定多个，渲染模板时无需指定扩展名，将根据配置寻找对应的文件。

##### Functions `map[string]interface{}`

扩展函数，参考：[FuncMap](https://godoc.org/html/template#FuncMap)

#### 扩展语法

##### 加载模板片段 

```
{{ template "share/footer" }}
```

### pongo2

`https://github.com/go-baa/pongo2`

pongo2 是 一个将 [Pongo2](https://github.com/flosch/pongo2) 模板引擎应用到 baa的助手库。

#### 使用

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

#### 配置 `pongo2.Options`

##### Baa `*baa.Baa`

render 需要传递 baa 实例

##### Root `string`

模板目录，指定模板目录，`c.HTML` 渲染模板时将从该目录下查找目录，不需要在渲染是指定目录名。

##### Extensions `[]string`

模板文件扩展名，可以指定多个，渲染模板时无需指定扩展名，将根据配置寻找对应的文件。

##### Functions `map[string]interface{}`

扩展函数，参考：[FuncMap](https://godoc.org/html/template#FuncMap)

##### Context `map[string]interface{}`

预置数据，模板变量

#### 扩展语法

##### 输出变量

```
{{ name }}
```

##### 加载模板片段

```
{% include "path/to/tpl.html" %}
```

带参数加载

```
{% include "relative/path/to/tpl.html" with foo=var %}
{% include "relative/path/to/tpl.html" with foo="bar" %}
```

> 嵌入的模板接收的参数将作为 `string` 类型。

##### 条件语句

```
{% if vara %}
{% elif varb %}
{% else %}
{% endif %}
```

##### 循环语句

```
{% for item in items %}
{{ forloop.Counter }} {{ forloop.Counter0 }} {{ forloop.First }} {{ forloop.Last }} {{ forloop.Revcounter }} {{ forloop.Revcounter0 }}
{{ item }}
{% endfor %}
```

##### 内置过滤器

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

#### extends / block / macro and so on ...

更多内容，参见 [django](https://docs.djangoproject.com/en/dev/ref/templates/language/).
