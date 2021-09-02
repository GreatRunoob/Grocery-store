/*
即时通信系统主程序源代码
*/

package main

import "flag"

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "Set server ip(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "Set server port(default:8888)")

	// 命令行格式：./server -ip 127.0.0.1 -port 8888
	// init()根据上述格式进行命令行参数解析
}

func main() {
	flag.Parse()

	// 创建并初始化一个服务端，设置监听本地主机8888端口
	server := NewServer(serverIp, serverPort)

	// 启动服务端
	server.Start()
}
