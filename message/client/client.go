/*
本源代码包含即时通信系统所需的代码
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp    string
	ServerPort  int
	conn        net.Conn
	choice      int           // 当前客户端通信模式
	inputReader *bufio.Reader // 支持cli模式带特殊字符读取
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:    serverIp,
		ServerPort:  serverPort,
		choice:      999, // 如果不指定初始值，默认初始化为0，则会直接退出通信业务
		inputReader: bufio.NewReader(os.Stdin),
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	client.conn = conn

	return client
}

// 客户端功能菜单
func (this *Client) menu() bool {
	var choice int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Println("请输入菜单选项:")
	fmt.Scanln(&choice) // 等待用户输入选项时阻塞

	if choice >= 0 && choice <= 3 {
		this.choice = choice
		return true
	} else {
		fmt.Println("请输入合法的选项！")
		return false
	}
}

func (this *Client) PublicChat() {
	var chatMsg string

	fmt.Println("请输入广播内容(输入q退出):")
	if err := this.Input(&chatMsg); err != nil {
		return
	}

	for chatMsg != "q" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			if ok := this.SendMsg(sendMsg); !ok {
				fmt.Println("消息发送失败！")
				break
			}

		} else {
			fmt.Println("广播内容不能为空！")
		}

		fmt.Println("请输入广播内容(输入q退出):")
		if err := this.Input(&chatMsg); err != nil {
			return
		}
	}
}

// 查询在线用户API
func (this *Client) OnlineUsers() bool {
	sendMsg := "who\n"
	isOk := this.SendMsg(sendMsg)

	return isOk
}

// 私聊模式
func (this *Client) PrivateChat() {
	// 首先反馈给用户，当前服务端所有在线用户信息(who)
	if ok := this.OnlineUsers(); !ok {
		fmt.Println("在线用户列表获取失败！")
		return
	}

	var remoteName string
	var chatMsg string

	fmt.Println("请输入私聊对象的用户名(输入q退出):")
	fmt.Scanln(&remoteName)

	// 接着让用户选择私聊对象，完成私聊业务
	for remoteName != "q" {

		fmt.Println("请输入私聊消息内容(输入q退出):")
		if err := this.Input(&chatMsg); err != nil {
			return
		}

		for chatMsg != "q" {
			if len(chatMsg) != 0 {
				sendMsg := "to:" + remoteName + ":" + chatMsg + "\n"
				if ok := this.SendMsg(sendMsg); !ok {
					fmt.Println("消息发送失败！")
					break
				}

			} else {
				fmt.Println("私聊内容不能为空！")
			}

			fmt.Println("请输入私聊消息内容(输入q退出):")
			if err := this.Input(&chatMsg); err != nil {
				return
			}
		}

		fmt.Println("请输入私聊对象的用户名(输入q退出):")
		fmt.Scanln(&remoteName)
	}
}

func (this *Client) UpdateName() bool {
	var newName string

	fmt.Println("请输入新用户名:")
	fmt.Scanln(&newName)

	sendMsg := "rename:" + newName + "\n"
	isOk := this.SendMsg(sendMsg)

	return isOk
}

// 对客户端向服务端发送消息做封装
func (this *Client) SendMsg(msg string) bool {
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

// 自定义用户长串消息读取
func (this *Client) Input(pStr *string) error {
	input, err := this.inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("inputReader.ReadString err:", err)
	} else {
		*pStr = input[:len(input)-1] // 删去末尾多余的换行符
	}

	return err
}

// 处理server回应的消息
func (this *Client) DealResponse() {
	// 一旦conn有数据，就直接copy到stdin上，永久阻塞等待消息
	io.Copy(os.Stdout, this.conn)
}

// 客户端通信业务运行逻辑
func (this *Client) Run() {
	for this.choice != 0 {
		for this.menu() != true {
		}

		// 根据不同的模式处理不同的通信业务
		switch this.choice {
		case 1:
			// 公聊模式
			this.PublicChat()
			break
		case 2:
			// 私聊模式
			this.PrivateChat()
			break
		case 3:
			// 更新用户名
			this.UpdateName()
			break
		}
	}
}
