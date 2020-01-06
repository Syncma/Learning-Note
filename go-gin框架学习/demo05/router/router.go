package router

import (
	"apiserver/handler/api"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// load middlewares, routes, handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middlewares
	//在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic
	//这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(gin.Recovery())

	//强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	//浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Options)
	//一些安全设置
	g.Use(middleware.Secure)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	// The health check handlers
	svcd := g.Group("/api")
	{
		svcd.GET("/health", api.HealthCheck)
	}

	return g
}
