# GO gin框架学习笔记-Day1

[toc]


## Gin是什么

Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。 

它是一个类似于 martini 但拥有更好性能的 API 框架, 由于 httprouter，速度提高了近 40 倍。

## 开发环境

环境介绍：
```
go版本： v1.13.1
编辑器：vscode
采用 Go Modules 进行管理
```

### 题外话-go mod
  go mod 是Golang 1.11 版本引入的官方包（package）依赖管理工具，用于解决之前没有地方记录依赖包具体版本的问题，方便依赖包的管理。

之前Golang 主要依靠vendor和GOPATH来管理依赖库，vendor相对主流，

但现在官方更提倡go mod


#### 配置
下载官方包1.11(及其以上版本将会自动支持gomod) 默认GO111MODULE=auto(auto是指如果在gopath下不启用mod)

Golang 提供一个环境变量 GO111MODULE 来设置是否使用mod，它有3个可选值，分别是off, on, auto（默认值），具体含义如下：

>off: GOPATH mode，查找vendor和GOPATH目录

>on：module-aware mode，使用 go module，忽略GOPATH目录

>auto：如果当前目录不在$GOPATH 并且 当前目录（或者父目录）下有go.mod文件，则使用 GO111MODULE， 否则仍旧使用 GOPATH mode。


默认是这样的：
```
GO111MODULE=""
```


修改 GO111MODULE 的值的语句是：```set GO111MODULE=on ```

在使用模块的时候， GOPATH 是无意义的，不过它还是会把下载的依赖储存在 GOPATH/src/mod 中，也会把 go install 的结果放在 GOPATH/bin（如果 GOBIN 不存在的话）


go mod命令：
>go mod download 下载模块到本地缓存，缓存路径是 $GOPATH/pkg/mod/cache

>go mod edit 是提供了命令版编辑 go.mod 的功能

>例如 go mod edit -fmt go.mod 会格式化 go.mod

>go mod graph 把模块之间的依赖图显示出来

>go mod init 初始化模块（例如把原本dep管理的依赖关系转换过来）

>go mod tidy 增加缺失的包，移除没用的包

>go mod vendor 把依赖拷贝到 vendor/ 目录下

>go mod verify 确认依赖关系

>go mod why 解释为什么需要包和模块


* 注意有几个坑的地方：

1.go mod 命令在 `$GOPATH` 里默认是执行不了的，因为 GO111MODULE 的默认值是 auto。

默认在$GOPATH 里是不会执行， 如果一定要强制执行，就设置环境变量为 on。

2.go mod init 在没有接module名字的时候是执行不了的，会报错 go: cannot determine module path for source directory



#### 例子
1.在`$GOPATH` 目录之外新建一个目录，并使用go mod init 初始化生成go.mod 文件

```
go mod init name (name可以随便写)
```


go.mod文件一旦创建后，它的内容将会被go toolchain全面掌控。

go toolchain会在各类命令执行时，比如go get、go build、go mod等修改和维护go.mod文件。


2.创建main.go
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

```

执行：
```go
go run main.go
```

3.使用go run 命令运行，发现go mod会自动查找依赖并自动下载

如果出现timeout，解决办法：**`使用七牛国内代理`**：

[七牛国内代理](https://github.com/goproxy/goproxy.cn)



使用方法：

```
go env -w GOPROXY=https://goproxy.cn,direct
```

然后再次执行 go run 操作

go module 安装 package 的原則是先拉最新的 release tag，若无tag则拉最新的commit，详见 Modules官方介绍。

go 会自动生成一个 go.sum 文件来记录 dependency tree


#### 特点

使用 Go Modules 后，不在需要以下内容：

>1.不用再定义 GOPATH （这里指的是 go build 、 go install 等等 go 命令。IDE 插件目前还是需要 GOPATH）

>2.工程目录放置，不再需要 src 目录下 （同上情况）

>3.不再需要 vendor 机制以及其他第 3 方 dep 工具

>4.工程内不再有依赖库代码。

>5.使用 Go Modules 后，理论上：
	* 代码可以随意放置
	* 执行 Go 命令，不再需要指定 GOPATH

模块会自动下载到`$GOPATH/pkg`目录：

[Go mod用法指南](http://wjp2013.github.io/go/go-module/)




## 设计阶段

### REST VS RPC

普遍采用的做法是，内部系统之间调用用 RPC，对外用 REST，因为内部系统之间可能调用很频繁，需要 RPC 的高性能支撑。对外用 REST 更易理解，更通用些。


Go 语言中常用的 API 风格是 RPC 和 REST，常用的媒体类型是 JSON、XML 和 Protobuf。

在 Go API 开发中常用的组合是 gRPC + Protobuf 和 REST + JSON。

这里使用**REST+JSON**方式进行开发


###  GOPATH VS GOROOT

1.GOPATH的作用是告诉Go 命令和其他相关工具，在那里去找到安装在你系统上的Go包。
 作为编译后二进制的存放目的地和 import 包时的搜索路径


2.GOROOT是指安装支撑Go运行环境的目录


OK,  下面会分几个章节分别讲解







