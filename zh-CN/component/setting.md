# Baa setting

`https://github.com/go-baa/setting`

应用配置和配置文件管理

> 注意: 配置文件被硬编码为 `./conf/app.ini`，必须放在这个路径

## 示例

```
// conf/app.ini
[default]
# app
app.name = baaBlog
app.version = 0.1
app.url = ""
debug = false

# http
http.address = 0.0.0.0
http.port = 80
http.access_open = off

# output log to os.Stderr
log.file = os.Stderr
# 0 off, 1 fatal, 2 panic, 5 error, 6 warn, 10 info, 11 debug
log.level = 11

# development mode overwrite default config
[development]
debug = true

# production mode overwrite default config
[production]
debug = false
```

## 使用

```
package main

import "github.com/go-baa/setting"

func main() {
    appName := setting.AppName
    httpAddress := setting.Config.MustString("http.address", "127.0.0.1")
    httpPort := setting.Config.MustInt("http.port", 1323)
    httpAccessOpen := setting.Config.MustBool("http.access_open", false)
}
``
