package routing

import "github.com/gin-gonic/gin"

func HtmlProcessor(request *gin.Context, paths []string, method string) (response gin.H, template string) {
	switch paths[1] {
	// 选择式路由
	case "test":
		return nil, "index.html"
	case "base":
		return nil, "BASE.html"
	case "docs":
		return nil, "docs.html"

	default:
		return nil, "login.html"
	}
}
