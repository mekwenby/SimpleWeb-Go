package routing

/*
集中返回式Html写法
Build Mek
time 20240223
*/

import (
	"SimpleWeb/databases"
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
		id := request.Query("id")
		imageFileList := databases.GetImageFileList(id)
		for i := range imageFileList {
			imageFileList[i].Filepath = strings.ReplaceAll(imageFileList[i].Filepath, "\\", "/")
		}
		return gin.H{"id": id, "list": imageFileList}, "view.html"
	default:
		return nil, "login.html"
	}
}
func HtmlDemo(path string) (template string) {
	return path + ".html"
}
