# Baa log

`https://github.com/go-baa/log`

一个增强的日志管理器实现。

相比标准库中的 `log` 增加了如下特性：

* 可配置日志级别，Info/Warn/Error/Debug
* 更多的日志方法
* 优化的日志写入
* 开箱即用的设置

## 安装

* 依赖 [setting](github.com/go-baa/setting) 包
* 可选配置文件：`conf/app.ini`

```
go get -u github.com/go-baa/setting
go get -u github.com/go-baa/log
```

## 使用

```
package main

import (
    "github.com/go-baa/log"
)

func main() {
    log.Println("xx")
    log.Debugln("xx")
}
```

## 配置

配置，依赖 `setting` 配置文件，请在配置文件中加入以下配置：

```
// conf/app.ini
[default]
# output log to os.Stderr or filepath
log.file = os.Stderr
# 0 off, 1 fatal, 2 panic, 5 error, 6 warn, 10 info, 11 debug
log.level = 11
```

`log.file` 指定日志输出路径，可以是具体的文件，也可以是 `os.Stderr` 或 `os.Stdout`
`log.level` 日志级别，默认是 `5`，级别越大输出的错误越详细
