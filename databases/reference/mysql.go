package databases

/*

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

// 在应用程序初始化阶段创建数据库引擎
func init() {
	Engine = createEngine()
}

// 创建数据库引擎
func createEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	return engine
}

username是你的MySQL用户名
password是你的MySQL密码
127.0.0.1:3306是MySQL服务器的地址和端口号
dbname是你要连接的数据库名
charset=utf8mb4是设置字符集，确保支持中文等特殊字符。

*/
