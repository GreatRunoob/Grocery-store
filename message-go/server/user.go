/*
本源代码定义了即时通讯系统有关用户部分的模型
包含对用户模型及其操作的封装，尽可能地与服务端业务分离
*/

package main

import (
	"net"
	"strings"
)

// 定义目前所需的用户属性字段
type User struct {
	Name string
	Addr string      // 用户上线ip
	Chan chan string // 用户通信频道，目前暂用于接收服务端的消息

	// 以下为私有属性

	conn   net.Conn // 绑定用户自身的socket连接，用于与服务端之间交互
	server *Server  // 将用户自身的socket连接与对应的服务端绑定
}

//创建一个用户的API，传入用户连接服务端的socket以及对应的服务端引用
func NewUser(conn net.Conn, server *Server) *User {
	// 获取用户上线ip信息
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr, //目前暂时将用户上线ip作为用户名称字段，确保唯一性
		Addr:   userAddr,
		Chan:   make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听当前user channel消息的goroutine，用于及时将服务端消息广播给用户
	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	// 新用户上线，将自身加入到对应server维护的onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this // OnlineMap[string]*User
	this.server.mapLock.Unlock()

	// 向当前服务端上所有的在线用户广播新用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	// 用户下线，将用户从当前服务端的OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户向服务端发送广播消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {

		// 查询当前在线用户信息，消息格式：who
		this.server.mapLock.Lock()

		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}

		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename:" {

		// 在线用户改名，消息格式：rename:newname
		newName := strings.Split(msg, ":")[1]

		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名已被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已更新用户名:" + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to:" {

		// 给在线用户发送私聊消息，消息格式：to:username:message
		msgSlice := strings.Split(msg, ":")

		remoteName := msgSlice[1] // 获取消息接收方的用户名
		if remoteName == "" || len(msgSlice) < 3 {
			this.SendMsg("私聊消息格式不正确，请使用\"to:username:msg\"发送私聊消息\n")
			return
		}

		// 解析接收对象、消息内容
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户当前不在线或不存在\n")
			return
		}

		// 还原并检查私聊消息格式
		msg = strings.Join(msgSlice[2:], ":")
		if msg == "" {
			this.SendMsg("无消息内容，请重新编辑后发送\n")
			return
		}

		// 向目标用户发送私聊消息
		remoteUser.SendMsg(this.Name + "对您说:" + msg + "\n")

	} else {
		// 普通的广播消息，所有在线用户都能接收到
		this.server.BroadCast(this, msg)
	}
}

//监听当前user channel，一旦有消息就直接广播给用户
func (this *User) ListenMessage() {
	// 通过不断轮询user channel的方式
	// 其实也算不上轮询，如果没有消息则会被阻塞
	for {
		// 从user channel中获取服务端广播消息反馈给用户
		msg, ok := <-this.Chan

		// 检测用户通信频道是否已被主动关闭
		if !ok {
			break
		}

		// 格式化服务端的广播消息，约定默认末尾缺少换行符
		// 针对其他需要广播给用户的消息也要注意格式的问题
		this.conn.Write([]byte(msg + "\n"))
	}
}
