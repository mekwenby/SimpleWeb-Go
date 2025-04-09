package routing

/*
集中返回式Api写法
Build Mek
time 20240223
*/

import (
	"SimpleWeb/databases"
	"fmt"
	"github.com/gin-gonic/gin"
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
	case "NewUser": // 创建用户
		if method == "POST" {
			// 从POST表单获取参数
			name := request.PostForm("name")
			passwd := request.PostForm("passwd")
			user, err := databases.CreateUser(name, passwd)
			if err != nil {
				return gin.H{"status": false, "err": err}
			}
			return gin.H{
				"status": true, "user": user,
			}
		}
		return gin.H{"status": false, "method": method}
	case "UserOn": // 用户登录
		// 从POST表单获取参数
		if method == "POST" {
			name := request.PostForm("name")
			passwd := request.PostForm("passwd")
			user, err := databases.VerifyPassword(name, passwd)
			if err != nil {
				return gin.H{"status": false, "err": err}
			}
			request.SetCookie("token", user.Token, 60*60*24*3, "/", "", false, true)
			return gin.H{
				"status": true, "user": user,
			}
		}
		return gin.H{"status": false, "method": method}
	case "TokenOn": // 验证Token
		token, _ := request.Cookie("token")
		fmt.Println(token)
		user, err := databases.VerifyToken(token)
		if err != nil {
			return gin.H{"status": false, "err": err}
		}
		return gin.H{
			"status": true, "user": user,
		}
	// 路由匹配失败时
	default:
		return gin.H{
			"path":    paths,                                      // 获取访问的完整路由
			"host":    request.Request.Host,                       // 获取用户的主机地址
			"message": "没有该路由!,There is no such route available!", // 自定义消息
			"method":  method,
		}
	}

}
