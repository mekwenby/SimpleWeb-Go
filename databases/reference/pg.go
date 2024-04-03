package databases

/*
import (
"github.com/go-xorm/xorm"
_ "github.com/lib/pq"
)

var Engine *xorm.Engine

// 在应用程序初始化阶段创建数据库引擎
func init() {
	Engine = createEngine()
}

// 创建数据库引擎
func createEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", "user=username password=password dbname=dbname host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return engine
}

user是你的PostgreSQL用户名
password是你的PostgreSQL密码
dbname是你要连接的数据库名
host是PostgreSQL服务器的地址
port是PostgreSQL服务器的端口号，默认为5432
sslmode=disable是关闭SSL模式，你可以根据你的实际需求设置其他模式。

*/
