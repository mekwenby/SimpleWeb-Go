package routing

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Majestic(request *gin.Context, paths []string, method string) {
	switch paths[1] {
	case "index":
		request.HTML(http.StatusOK, "MajesticAdmin.html", nil)

	}

}
