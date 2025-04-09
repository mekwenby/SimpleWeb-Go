package databases

/*
数据库引擎
链接数据库服务使用
具体使用方法参考reference示例
*/

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var Engine *xorm.Engine

func init() {
	// 在应用程序初始化阶段创建数据库引擎
	Engine = CreateEngine()
}
func CreateEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "sqlite3.db")
	if err != nil {
		panic(err)
	}
	return engine
}
