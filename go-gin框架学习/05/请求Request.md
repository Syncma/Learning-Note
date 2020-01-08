# 请求Request

[toc]

## 请求头

Gin获取Http请求头Header和Body

一个HTTP报文由3部分组成，分别是:

1. 起始行(start line)

2. 首部(header)

3. 主体(body)


```go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Printf("launch Gin")

	r := gin.Default()
	r.GET("/get", HandleGet)
	r.POST("/getall", HandleGetAllData)

	//如果使用浏览器调试，那么响应Get方法
	//r.GET("/getall",HandleGetAllData)
	r.Run(":9000")

}

func HandleGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"receive": "65536",
	})

}

func HandleGetAllData(c *gin.Context) {
	//log.Print("handle log")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	fmt.Println("header \r\n", c.Request.Header)

	c.JSON(200, gin.H{
		"receive": "1024",
	})

}

```
　　
## 请求参数


## Cookies

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/setcookie", SetHandler)
	router.GET("/getcookie", GetHandler)
	router.GET("/delcookie", DelHandler)

	router.Run()
}

func SetHandler(c *gin.Context) {
	c.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	c.String(200, "cookie演示")
}

func GetHandler(c *gin.Context) {
	// 根据cookie名字读取cookie值
	data, err := c.Cookie("site_cookie")
	if err != nil {
		// 直接返回cookie值
		c.String(200, data)
		return
	}
	c.String(200, "not found!")
}

func DelHandler(c *gin.Context) {
	// 设置cookie  MaxAge设置为-1，表示删除cookie
	c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
	c.String(200, "删除cookie演示")
}

```


SetCookie参数说明：

| 参数名    |     类型|   说明  |
| :-------- | --------:| :------: |
| name    |   string |  cookie名字  |
| value	  |   string |  cookie值  |
| maxAge  |   int	|  有效时间，单位是秒，MaxAge=0 忽略MaxAge属性，MaxAge<0 相当于删除cookie, 通常可以设置-1代表删除，MaxAge>0 多少秒后cookie失效  |
| path	    |   string |  cookie路径  |
| domain	|   string |  cookie作用域  |
| secure	|   bool	|  Secure=true，那么这个cookie只能用https协议发送给服务器  |
| httpOnly	|   bool| 设置HttpOnly=true的cookie不能被js获取到  |


		
		


## 上传文件

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {

		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})
	router.Run(":8080")
}

```