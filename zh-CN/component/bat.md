# Baa bat

`https://github.com/go-baa/bat`

bat 是一个帮助快速开发 baa 程序的小工具。

代码fork自 [beego](https://github.com/astaxie/beego/) 的 [bee](https://github.com/beego/bee)。

## 使用

### 安装

```
go get -u github.com/go-baa/bat
```

## 运行baa程序

```
bat run [-x=.go -x=.ini] [-a=../model] [-e=Godeps -e=folderToExclude] [-tags=goBuildTags]
```

默认情况下监控运行目录下的 `.go`文件，发生变化就会重新编译并运行。

### godeps

bat 默认开启了 `godeps` 支持，如果项目下存在 `Godeps`目录，每次重新构建都会重建 `Godeps`。

可以通过以下参数关闭该特性：

```
bat run -godeps=false
```
