package engine

import (
	"SimpleWeb/routing"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Engine = gin.Default()

type RouteMountTable map[string]func()

func init() {
	// 创建gin引擎
	// 设置静态文件路径
	Engine.Use(static.Serve("/static", static.LocalFile("./static", true)))

	// 设置模板路径
	Engine.LoadHTMLGlob("templates/**/*")

	// 定义首页路由函数
	Engine.GET("/", Index)

	// 定义首页图标
	Engine.GET("/favicon.ico", func(c *gin.Context) {
		c.File("static/img/favicon.ico")
	})

	// 拦截404
	Engine.NoRoute(func(c *gin.Context) {
		// 返回404状态码和一条消息
		c.JSON(http.StatusNotFound, gin.H{"request_ip": c.ClientIP(),
			"path":       c.Request.URL.Path,
			"paths":      strings.Split(c.Request.URL.Path, "/")[1:],
			"method":     c.Request.Method,
			"message":    "NoRoute",
			"statuscode": 404,
		})
	})

	// 其他路由统一路径入口
	Engine.Any("/api/*path", Unified)
	Engine.Any("/html/*path", Unified)
	Engine.Any("/Api/*path", Unified)
	Engine.Any("/Html/*path", Unified)

}

func Unified(c *gin.Context) {
	// 请求路径
	path := c.Request.URL.Path
	// 请求主体对象
	request := c.Request
	// 请求方式
	method := request.Method
	// 请求路径列表
	parts := strings.Split(path, "/")[1:]

	switch parts[0] {
	case "Api":
		fallthrough
	case "api":
		c.JSON(http.StatusOK, routing.ApiProcessor(c, parts, method))
	case "Html":
		fallthrough
	case "html":
		response, template := routing.HtmlProcessor(c, parts, method)
		c.HTML(http.StatusOK, template, response)
	}

}

// Index 定义首页路由
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Default(c *gin.Context) {

	// 请求路径
	path := c.Request.URL.Path
	// 请求主体对象
	request := c.Request
	// 请求方式
	method := request.Method
	// 请求路径列表
	parts := strings.Split(path, "/")[1:]
	c.JSON(http.StatusOK, gin.H{

		"request_ip": request.Host,
		"path":       path,
		"paths":      parts,
		"method":     method,
	})
}

// Redirect 重定向
func Redirect(c *gin.Context, path string) {
	c.Redirect(http.StatusMovedPermanently, path)
}
