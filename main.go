package main

import (
	"SimpleWeb/engine"
	"SimpleWeb/tools"
	"flag"
	"strconv"
)
import "fmt"

func main() {
	startPort := flag.Int("port", 8086, "启动端口")
	flag.Parse()
	// databases.SyncTable()
	tools.Logo_Slabt("ImageBrowser")
	// "8086 启动端口号"
	fmt.Printf("http://127.0.0.1:%v", startPort)
	hostAddr := ":" + strconv.Itoa(*startPort)
	err := engine.Engine.Run(hostAddr)
	if err != nil {
		fmt.Print("Gin Server Run mistake. \nGin引擎运行失败!\n")
	}
}
