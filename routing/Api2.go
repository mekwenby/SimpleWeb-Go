package routing

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Api2(request *gin.Context, paths []string, method string) {
	// 动态返回式路由处理
	request.JSON(http.StatusOK, "Api2")
}
