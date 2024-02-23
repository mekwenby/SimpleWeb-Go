package engine

/*
常见公共方法拓展
Build Mek
time 20240223

	登录
	注销
	文件上传
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func Login(c *gin.Context) {

	if true { // 成功
		r := gin.H{
			"state": true,
		}
		c.SetCookie("token", "998", 36000, "/", "", false, true)
		c.JSON(http.StatusOK, r)
	} else { // 失败
		r := gin.H{
			"state": false,
		}
		c.JSON(http.StatusOK, r)

	}

}

func Logout(c *gin.Context) {
	// 删除与用户相关的Cookie，例如名为"token"的Cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	// 返回注销成功的消息
	r := gin.H{
		"message": "Logout successful",
	}
	c.JSON(http.StatusOK, r)
}

func UploadFile(c *gin.Context) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将文件保存到指定的路径
	uploadDir := "./uploads"
	filePath := filepath.Join(uploadDir, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	// 返回文件上传成功的消息
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File '%s' uploaded successfully", file.Filename), "status": true})
}
