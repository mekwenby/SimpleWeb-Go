package routing

/*
集中返回式Html写法
Build Mek
time 20240223
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

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
	case "view":
		return gin.H{}, "view.html"
	default:
		templateName := strings.Join(paths[1:], ".html")
		fmt.Println(templateName)
		return nil, templateName
	}
}
