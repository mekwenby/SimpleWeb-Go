package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo() {
	r := gin.Default()

	r.GET("/SteCookie", func(context *gin.Context) {
		domain := context.Request.Host
		context.SetCookie("token", "8090", 65432, "/", domain, false, true)
		context.String(http.StatusOK, "Set Cookie...")
	})

	r.GET("/delete-cookie", func(c *gin.Context) {
		// 将Cookie的过期时间设置为一个过去的时间点，以删除Cookie
		c.SetCookie("example-cookie", "", -1, "/", "localhost", false, true)

		// 返回一个消息
		c.String(http.StatusOK, "Cookie has been deleted")
	})
}
