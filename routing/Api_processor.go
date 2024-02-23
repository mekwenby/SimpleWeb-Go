package routing

import "github.com/gin-gonic/gin"

func ApiProcessor(request *gin.Context, paths []string, method string) (response gin.H) {
	switch paths[1] {
	default:
		return gin.H{
			"path": "api",
		}
	}

}
