package engine

import (
	"github.com/gin-gonic/gin"
	"net/http" // 导入net/http包
)

func Get_Cookies(request *gin.Context) []*http.Cookie {
	return request.Request.Cookies()
}
