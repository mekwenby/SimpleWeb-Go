package routing

/*
集中返回式Html写法
Build Mek
time 20240223
*/

import "github.com/gin-gonic/gin"

// HtmlProcessor 静态模板路由
func HtmlProcessor(request *gin.Context, paths []string, method string) (response gin.H, template string) {
	/*
		request:	请求主体
		paths:		请求路径列表
		method:		请求方式

		response:	返回JSON对象
		template:	模板名称
	*/
	switch paths[1] {
	// 选择式路由
	case "test":
		return nil, "index.html"
	case "base":
		return nil, "BASE.html"
	case "docs":
		return nil, "docs.html"
	case "login":
		return nil, "User_login.html"
	default:
		return nil, "login.html"
	}
}
