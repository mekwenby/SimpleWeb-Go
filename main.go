package main

import "SimpleWeb/engine"
import "fmt"

func main() {

	// "8081 启动端口号"
	err := engine.Engine.Run(":8081")
	if err != nil {
		fmt.Print("Gin Server Run mistake. \n引擎运行失败!\n")
	}
}
