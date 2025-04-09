package main

import (
	"SimpleWeb/databases"
	"SimpleWeb/engine"
	"SimpleWeb/tools"
	"flag"
	"strconv"
)
import "fmt"

func main() {
	startPort := flag.Int("port", 8086, "启动端口")
	startMode := flag.String("mode", "Server", "启动模式")
	flag.Parse()

	// 创建数据表
	if *startMode == "sync" {
		fmt.Println("同步结构体到数据库...")
		databases.SyncTable()
	}
	// 打印Logo
	tools.Logo_Slabt("SimpleWeb")
	// "8086 启动端口号"
	fmt.Printf("%v:http://127.0.0.1:%v\n", databases.GetCurrentTime(), *startPort)
	hostAddr := ":" + strconv.Itoa(*startPort)
	err := engine.Engine.Run(hostAddr)
	if err != nil {
		fmt.Print("Gin Server Run mistake. \nGin引擎运行失败!\n")
	}
}
