# Baa 工程化

所谓工程化，不是baa的功能，而是在使用baa的过程中总结出的一种姿势，姑且称之为：最佳实践。

按照一般MVC的开发规范，一个程序通常会有 数据模型、模板视图、业务控制，辅助的还要有 前端资源文件，应用配置文件，可能还要记录日志文件。

## 目录结构

```
project
  |-- conf
      |-- app.ini
  |-- controller
      |-- index.go
      |-- article.go
      |-- user.go
  |-- data
      |-- log
  |-- model
      |-- base
          |-- base.go
          |-- cache.go
      |-- base.go
      |-- article.go
      |-- user.go
  |-- module
      |-- util
          |-- util.go
      |-- template
          |-- template.go
  |-- public
      |-- assets
          |-- css
          |-- images
          |-- js
      |-- upload
      |-- robots.txt
  |-- router
      |-- init.go
      |-- router.go
  |-- template
      |-- share
          |-- header.html
          |-- footer.html
      |-- article
          |-- index.html
          |-- show.html
      |-- user
          |-- show.html
      |-- index.html
  |-- main.go
  |-- README.md
```

结构说明

| 路径                         | 说明          | 备注   |
|-----------------------------|--------------|--------|
| conf                        | 配置文件目录   | --     |
| conf/app.ini                | 应用配置文件   | [setting](https://github.com/go-baa/setting) 配置库要求的配置文件路由
| controller                  | 业务控制器目录 | --     |
| controller/*.go             | 具体控制器     | 建议每个功能一个控制器文件 |
| data                        | 数据目录      | --     |
| data/log                    | 日志目录      | 建议路径，可在配置文件中指定 [log](https://github.com/go-baa/log) 输出路径 |
| model                       | 数据模型目录   | --     |
| model/base                  | 数据模型基类   | 提供对于数据库连接，缓存连接等基础操作 |
| model/base.go               | 模型基类      | 导入 `model/base` 初始化数据库，是其他模型的基础 |
| model/article.go            | 业务模型      | 具体的业务模型，建议每个表对应一个模型文件，命名和表名一致 |
| module                      | 扩展功能模块   | --     |
| module/util                 | 助手模块      | 一些常用的功能函数，文件操作，URL操作，加解密等 |
| module/util/util.go         | 助手函数      | 常用函数库 |
| module/template             | 模板扩展      | --     |
| module/template/template.go | 模板函数库    | 结合 [render](https://github.com/go-baa/render) 模板引擎，可以扩展模板函数 |
| public                      | 静态资源目录   | --     |
| public/assets               | 前端资源目录   | --     |
| public/assets/css,images,js | 前端文件目录   | --     |
| public/uplaod               | 上传文件目录   | --     |
| public/robots.txt           | 静态文件      | 其他静态文件，可以放在资源目录下 |
| router                      | 路由设定目录   | --     |
| router/init.go              | baa初始化     | 初始化baa，加载中间件，模板组件，缓存组件等 |
| router/router.go            | 路由配置      | 独立了路由配置在一个文件中，结构更清晰 |
| template                    | 模板目录      | --      |
| template/share              | 共享目录      | 存储共享的模板片段 |
| template/article            | 业务模板      | 具体的业务模板，建议和控制一一对应，每个控制一个目录，每个方法一个文件 |
| template/index.html         | 首页模板      | 应用的首页文件 |
| main.go                     | 应用入口      | --     |
| README.md                   | 应用说明      | --     |

## 控制器

控制器中按业务划分成了不同的文件，不同的操作还应该有不同的方法对应，在实现上有两种考虑：

- 一个控制器中所有方法都是函数，使用控制器的名字作为函数名前置防止多个控制中的命名冲突。
- 将一个控制器视为一个类，所有方法都是类的方法，虽然Go中没有明确的类，但也可以实现面向对象编程。

两种声音都有支持，你可以根据自己喜欢来做，我们选择了第二种姿势，看起来更舒服一些。

最终，一个控制文件可能是这样的：

```
// api/controller/index.go

package controller

import (
	"github.com/go-baa/example/api/model"
	"github.com/go-baa/log"
	"gopkg.in/baa.v1"
)

type index struct{}

// IndexController ...
var IndexController = index{}

// Index list articles
func (index) Index(c *baa.Context) {
	page := c.ParamInt("page")
	pagesize := 10

	rows, total, err := model.ArticleModel.Search(page, pagesize)
	if err != nil {
		output(c, 1, err.Error(), nil)
		return
	}

	log.Debugf("rows: %#v, total: %d\n", rows, total)

	output(c, 0, "", map[string]interface{}{
		"total": total,
		"items": rows,
	})
}

....

```

> 该文件来自示例程序 [api](http://github.com/go-baa/example/tree/master/api)

为了实现面向对象，创建了一个空的结构体作为方法的承载，所有方法都注册给这个结构体。

**需要解释的一句是，为什么还要声明一个 `IndexController` 呢？**

路由注册时需要将每一个URL对应到具体的方法上来，结构体的方法是不能直接用的，需要先声明一个结构体实例才能使用。

在哪儿声明呢？一个是路由注册的时候，一个是控制器定义的时候，我们选择了在控制器定义的时候声明，作为控制器开发的一个规范，路由定义时引入包就可以用了。

## 数据模型

baa本身不提供数据模型的处理，在 [api](http://github.com/go-baa/example/tree/master/api) 示例中使用的是 [grom](http://jinzhu.me/gorm/) 来操作MySQL。

[xorm](http://xorm.io/) 和 [grom](http://jinzhu.me/gorm/) 有什么区别呢？论功能 `xorm` 可能更强大一些，我们觉得 `grom` 使用更舒服一些。

虽然他们都做了很好的封装，但一个项目毕竟还要配置数据库信息，数据库连接，还要各种包调用，显然我们还是要做个简单的封装才好。

具体的代码不列出，请参考 [api/model](http://github.com/go-baa/example/tree/master/api/model) 中的base处理。

在这个基础上，一个数据模型可能长这个样子：

```
// api/model/user.go

package model

// User user data scheme
type User struct {
	ID    int    `json:"id" gorm:"primary_key; type:int(10) UNSIGNED NOT NULL AUTO_INCREMENT;"`
	Name  string `json:"name" gorm:"type:varchar(50) NOT NULL DEFAULT '';"`
	Email string `json:"email" gorm:"type:varchar(100) NOT NULL DEFAULT '';"`
}

type userModel struct{}

// UserModel single model instance
var UserModel = new(userModel)

// Get find a user info
func (t *userModel) Get(id int) (*User, error) {
	row := new(User)
	err := db.Where("id = ?", id).First(row).Error
	return row, err
}

// Create create a user
func (t *userModel) Create(name, email string) (int, error) {
	row := new(User)
	row.Name = name
	row.Email = email
	err := db.Create(row).Error
	if err != nil {
		return 0, err
	}
	return row.ID, nil
}
```

> 该文件来自示例程序 [api](http://github.com/go-baa/example/tree/master/api)

基本思想和控制器是一样的，先按照表结构声明一个结构体。然后创建一个空结构体将模型的方法进行封装，最后声明了一个 `UserModel`使得在控制器中无需声明就可以直接使用模型。

**需要注意的是，在模型中每个方法的最后一个参数一定是`error`，表示操作是否出错，不要问为什么，规范，还是规范，这里讲的都是规范。**

## 配置文件

应用配置文件，只能是 `conf/app.ini`，这个由项目 [setting](https://github.com/go-baa/setting) 决定，为什么把路径写死了呢，为了省事，无论在哪儿引入包就能用，无需配置和传递。

更多的配置文件也建议放在 `conf` 目录中，自己去读取。

## 模板

模板最简单，按着结构放就好了。

模板的初始化和使用，参考：

* [模板渲染](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板渲染)
* [模板语法](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板语法)
* [模板接口](https://github.com/go-baa/doc/tree/master/zh-CN/context.md#模板接口)
* [render](https://github.com/go-baa/doc/tree/master/zh-CN/component/render.md)
* [pongo2](https://github.com/go-baa/doc/tree/master/zh-CN/component/pongo2.md)

项目示例，参考：

* [blog](https://github.com/go-baa/example/tree/master/blog)

## 静态资源

如果是一个API项目，可能没有静态资源，忽略就行。

一般的静态资源，放在那里就好了，然后注册静态资源目录：

```
// router/router.go

app.Static("/assets", "public/assets", false, nil)
app.StaticFile("/robots.txt", "public/robots.txt")
```

如果你的项目采用了前端构建的姿势，那么你就构建吧，和baa也没什么关系，也不影响，

就是建议把构建后的资源放置到 `public`下面，比如：`public/assets/build` 然后注册静态目录，开发过程中的文件不建议放在 `public`下，因为是不可访问资源。

## 打包发布

Go程序的一个好处就是，`go build`然后生成一个二进制文件，Copy到服务上就行了。

不过需要注意的是，按照以上介绍的姿势，你还要带上 `配置文件`，`模板`，`静态资源`，最后运行的目录应该是这样的：

```
project
  |-- conf
      |-- app.ini
  |-- public
      |-- assets
          |-- build
      |-- css
      |-- images
      |-- js
      |-- robots.txt
  |-- template
      |- share
      |-- article
          |-- show.html
      |-- index.html
  |-- project // 二进制文件
```

至于你的发布姿势，是什么发布系统都没关系，要注意，打包的环境和运行的系统环境要一致，mac下编译出来的，linux可不一定能运行。

> PS：我们发布时采用 gitlab + jenkins 构建 Docker镜像的方式。

### 依赖管理

依赖管理的工具有很多，我们目前使用的是 [godep](https://github.com/tools/godep)，我们将产生的`Godeps`目录上传到了git中，确保构建时的环境一致。


