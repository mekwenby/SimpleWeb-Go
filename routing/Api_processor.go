package routing

/*
集中返回式Api写法
Build Mek
time 20240223
*/

import (
	"SimpleWeb/databases"
	"github.com/gin-gonic/gin"
	"strings"
)

// ApiProcessor API路由
func ApiProcessor(request *gin.Context, paths []string, method string) (response gin.H) {

	/*
		request:	请求主体
		paths:		请求路径列表
		method:		请求方式

		response:	返回JSON对象
	*/

	switch paths[1] {
	case "sync": // 重构数据库
		databases.CreateVersion()
		databases.BuildFileAssembly()
		info := databases.GetVersionInfo()
		return gin.H{
			"status":  true,
			"version": info.Version,
			"verify":  info.Verify,
		}

	case "version": // 获取版本信息
		info := databases.GetVersionInfo()
		return gin.H{
			"status":  true,
			"version": info.Version,
			"verify":  info.Verify,
		}

	case "copy":
		if method == "POST" {
			id := request.PostForm("id")
			imageFileList := databases.GetImageFileList(id)
			for i := range imageFileList {
				imageFileList[i].Filepath = strings.ReplaceAll(imageFileList[i].Filepath, "\\", "/")
			}
			go databases.AutoCopy(imageFileList)
			return gin.H{"method": method, "id": id, "list": imageFileList}
		} else {
			return gin.H{"Err": "请求方式不正确!"}
		}

	default:
		return gin.H{
			"path": "api",
		}
	}

}
