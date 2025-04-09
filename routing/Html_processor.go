package routing

/*
集中返回式Html写法
Build Mek
time 20240223
*/

import (
	"SimpleWeb/configs"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
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
	case "index":
		return nil, "index.html"

	default:
		// debug 模式下，返回所有信息
		if configs.Mode == "debug" {
			method := request.Request.Method
			path := request.Request.RequestURI
			// 获取请求头
			headers := request.Request.Header
			// 获取查询参数
			queryParams := request.Request.URL.Query()
			// 获取表单数据
			postForm := request.Request.PostForm
			// 获取请求体 (如需要，建议使用 gin 的 BindJSON 或 Bind 来解析请求体)
			// 此处示例仅显示原始的请求体内容（不适用于所有情况）
			var requestBody string
			if request.Request.Body != nil {
				buf := make([]byte, 1024)
				n, _ := request.Request.Body.Read(buf)
				requestBody = string(buf[:n])
			}
			// 获取当前时间戳
			requestTime := time.Now().Format(time.RFC3339)
			// 构建一个模板展示所有信息
			return gin.H{
				"method":      method,
				"path":        path,
				"headers":     headers,
				"queryParams": queryParams,
				"postForm":    postForm,
				"requestBody": requestBody,
				"requestTime": requestTime,
			}, "context_info.html"
		}
		// 生产模式下，正常返回模板
		// 解析模板名称
		templateName := strings.Join(paths[1:], ".html")
		return nil, templateName
	}
}
