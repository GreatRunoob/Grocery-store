/*
本源代码定义了即时通讯系统有关服务端模型
包含对服务端模型及其操作的封装
*/

package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 即时通信系统服务端模型
type Server struct {
	Ip   string // 服务端监听的ip地址
	Port int    // 服务端监听的端口号

	OnlineMap map[string]*User // 维护一个记录所有在线用户的map
	mapLock   sync.RWMutex     // 针对map操作的锁

	Message chan string // 服务端向在线用户广播消息的频道
}

// 创建并初始化一个服务端
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

//用于监听服务端Message广播消息，一旦有内容就发送给全部的在线用户
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将服务端广播msg发送给全部的在线用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.Chan <- msg // 向每个在线用户的通信频道中传入消息
		}
		this.mapLock.Unlock()
	}
}

// 服务端向在线用户广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	// 样例：[192.168.1.100]192.168.1.100:Hello everyone!

	this.Message <- sendMsg // 将消息传入服务端广播频道
}

// 服务端即时通信业务处理方法
func (this *Server) Handler(conn net.Conn) {

	fmt.Println("服务端链接建立成功")

	// 根据监听结果，新建一个用户对象
	user := NewUser(conn, this)

	user.Online() // 执行用户上线相关业务

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 开启一个goroutine，用于接收该用户客户端发送给服务端的消息
	go func() {
		buf := make([]byte, 4096) // 消息缓冲，大小自己决定
		for {
			msgLen, err := conn.Read(buf) // 当接收到客户端有消息事件时
			if msgLen == 0 {
				user.Offline() // 执行用户下线相关业务
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn.Read err:", err)
				return
			}

			// 获取用户消息并去除末尾的'\n'
			msg := string(buf[:msgLen-1])

			// 将获取到的用户消息向所有在线用户进行广播
			user.DoMessage(msg)

			// 用户的任意消息行为代表当前用户处于活跃状态
			isLive <- true
		}
	}()

	for {
		/*
			先给自己做个笔记，定时器这个地方会有点难理解
			有关time.After()方法的使用可以自行百度，当传入间隔时间参数后
			会定期向管道中传入时间信号，以达到定时的作用

			若用户一直处于活跃状态，便可通过isLive管道中的活跃信号激活
			select代码块，执行外isLive的case之后，同时第二个case部分，
			会重新生成一个新的定时器，重新开始计时
			从而达到更新用户连接超时定时器的作用
			可以这么说，只要select被外界激活一次，那么定时器就会被重置
			因为time.After是select代码块中的局部变量，
			本轮select内容执行结束后，对应的定时器将会被销毁
			等待select的下一次轮询重新生成定时器
		*/
		select {
		case <-isLive:
			// 当前User处于活跃状态，重置定时器
			// 主要作用是激活select更新定时器
			// 所以并无额外操作
		case <-time.After(time.Second * 180):
			// 已超时，将当前User连接强制关闭，超时时间可自行设置
			user.SendMsg("连接超时\n")

			/*
				这里有一个潜在的bug，当用户是因为超时而下线的情况
				如果不显式调用user.Offline()，理论上是没什么问题
				因为用户连接一旦被强制关闭，会激发一次消息事件
				由上面的协程go func()最终负责调用user.Offline()
				这也是上面代码设计的初衷。只不过作为一个子线程
				很有可能在本代码快退出时，未能及时调用user.Offline()
				导致用户在线列表得不到及时维护而出错，所以在此显式调用
				user.Offline()，同时也能让本代码段意图明显
				至于因二次维护在线用户列表是否会造成panic
				delete会帮忙处理好
			*/
			user.Offline()   // 用户下线，及时维护在线用户列表
			close(user.Chan) // 清理User的通信频道资源
			conn.Close()     // 关闭连接

			// 退出当前Handler
			return
		}
	}
}

// 启动服务端的接口
func (this *Server) Start() {
	// 监听指定ip的指定端口
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 程序退出时，关闭端口监听并释放资源
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessager()

	for {
		// 指定端口接收到新的用户连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// 处理用户通信业务
		go this.Handler(conn)

		/*
			服务端只需维护一个消息广播goroutine
			而每个在线用户需要单独维护自己的通信业务goroutine
		*/
	}
}
