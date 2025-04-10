package main

import (
	"SimpleWeb/configs"
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
	configs.Mode = *startMode
	// 创建数据表
	if *startMode == "sync" {
		fmt.Println("同步结构体到数据库...")
		databases.SyncTable()
		return
	}
	// 打印Logo
	tools.Logo_Slabt("SimpleWeb")
	// "8086 启动端口号"
	fmt.Printf("%v:http://127.0.0.1:%v\nRun Mode %v \n", databases.GetCurrentTime(), *startPort, *startMode)
	hostAddr := ":" + strconv.Itoa(*startPort)
	err := engine.Engine.Run(hostAddr)
	if err != nil {
		fmt.Print("Gin Server Run mistake. \nGin引擎运行失败!\n")
	}
}
