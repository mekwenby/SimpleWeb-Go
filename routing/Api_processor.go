package routing

/*
集中返回式Api写法
Build Mek
time 20240223
*/

import "github.com/gin-gonic/gin"

// ApiProcessor API路由
func ApiProcessor(request *gin.Context, paths []string, method string) (response gin.H) {

	/*
		request:	请求主体
		paths:		请求路径列表
		method:		请求方式

		response:	返回JSON对象
	*/

	switch paths[1] {
	default:
		return gin.H{
			"path": "api",
		}
	}

}
