package main

import (
	"flag"
	"fmt"
)

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "Set server ip(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "Set server port(default:8888)")

	// 命令行格式：./client -ip 127.0.0.1 -port 8888
	// init()根据上述格式进行命令行参数解析
}

func main() {
	flag.Parse() // 命令行解析

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("服务端连接失败！")
		return
	}

	// 单独处理server的回执消息
	go client.DealResponse()

	fmt.Println("服务端登录成功！")

	client.Run()
}
