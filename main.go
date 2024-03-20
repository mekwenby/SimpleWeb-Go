package main

import (
	"SimpleWeb/databases"
	"SimpleWeb/engine"
	"SimpleWeb/tools"
)
import "fmt"

func main() {
	err := databases.CreateUser("mek", "123", "EDF")
	if err != nil {
		return
	}

	tools.Logo_Slabt("SimpleWeb")
	// "8081 启动端口号"
	err = engine.Engine.Run(":8081")
	if err != nil {
		fmt.Print("Gin Server Run mistake. \n引擎运行失败!\n")
	}
}
