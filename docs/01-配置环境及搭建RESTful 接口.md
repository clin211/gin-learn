# Gin

Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。 它是一个类似于 [martini](https://github.com/go-martini/martini) 但拥有更好性能的 API 框架, 优于 [httprouter](https://github.com/julienschmidt/httprouter)，速度提高了近 40 倍

> Go 1.13 及以上版本

## 特性

- 快速；基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

- 支持中间件；传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

- Crash 处理；Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

- JSON 验证；Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

- 路由组；更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

- 错误管理；Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

- 内置渲染；Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

- 可扩展性

## 安装

要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作区。

1.下载并安装 gin：

```sh
$ go get -u github.com/gin-gonic/gin
```

2.将 gin 引入到代码中：

```go
import "github.com/gin-gonic/gin"
```

3.（可选）如果使用诸如 `http.StatusOK` 之类的常量，则需要引入 `net/http` 包：

```go
import "net/http"
```

- 创建你的项目文件夹并 `cd` 进去

	```sh
	$ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"
	```

- 拷贝一个初始模板到你的项目里

```sh
$ curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

- 运行你的项目

```sh
$ go run main.go
```

完整代码如下：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务，可以传一个监听端口的字符串 r.Run("9090")
	r.Run()
}
```

![image-20220103141047701](https://assets-website.oss-cn-chengdu.aliyuncs.com/notes/2022/01/03/14-10-52-1641190252-1641190252896-IlQZ83-14-10-48-1641190248-1641190248485-sA8Bb8-image-20220103141047701.png)

<img src="https://assets-website.oss-cn-chengdu.aliyuncs.com/notes/2022/01/03/14-17-15-1641190635-1641190635084-gohTTo-image-20220103141714912.png" alt="image-20220103141714912" style="zoom:50%;" />

程序虽然运行起来了，但是每次修改一点都要重新启动`go run main.go`，这样很影响我们开发体验，不仅影响开发体验，而且还是重复的劳动，程序员的就是为了减少重复劳动而实践的，所以说呢，经过一些调查发现了两个工具：

fresh [文档](https://github.com/gravityblast/fresh) 和 gin [文档](https://github.com/codegangsta/gin)

## fresh

### 安装

```sh
$ go get github.com/pilu/fresh
```

### 使用

在终端中进入项目的启动目录，然后运行命令：

```sh
$ fresh
```

> 项目启动后，会在项目的根目录中创建一个临时目录`tmp`，里面有个`runner-build`文件，这个文件就是来替代我们每次从新运行的复杂度，每次改动并保存后，都会重新运行项目

![image-20220103143150764](https://assets-website.oss-cn-chengdu.aliyuncs.com/notes/2022/01/03/14-31-50-1641191510-1641191510977-9b7MF8-image-20220103143150764.png)

## 搭建RESTful 接口

> CRUD操作

- GET

```go
r.GET("/user", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "message": "获取用户信息",
  })
})
```

- POST

```go
r.POST("/user", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "message": "创建一条用户信息",
  })
})
```

- PUT

```go
r.PUT("/user", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "message": "更新一条用户信息",
  })
})
```

- DELETE

```go
r.DELETE("/user", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "message": "删除一条用户信息",
  })
})
```

完整代码如下：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 路由
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "获取用户信息",
		})
	})
  
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "创建一条用户信息",
		})
	})

	r.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "更新一条用户信息",
		})
	})

	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "删除一条用户信息",
		})
	})

	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务，可以传一个监听端口的字符串 r.Run(":9090")
	r.Run(":9090")
}
```

