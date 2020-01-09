# 响应 Response

[toc]


## 响应头


## 附加cookie


## 字符串响应
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "some string")
	})
	router.Run(":8080")
}

```


## json响应

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		c.XML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		c.YAML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		c.JSON(http.StatusOK, msg)
		c.XML(http.StatusOK, msg)
		c.YAML(http.StatusOK, msg)
	})
	router.Run(":8080")
}

```


## 视图响应

LoadHTMLTemplates() 方法来加载模板文件

-例子待补充



## 文件响应

-待补充


## 重定向

-待补充



## 同步异步

goroutine 机制可以方便地实现异步处理



```go
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			// 注意使用只读上下文
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})
	//2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		// 注意可以使用原始上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

```