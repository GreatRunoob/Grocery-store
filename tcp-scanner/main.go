/*
TCP端口并发扫描器
模拟goroutine池实现并发扫描
支持命令行参数 -ip -from -to
*/

package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

var scanIp string
var fromPort int
var toPort int

func init() {
	flag.StringVar(&scanIp, "ip", "127.0.0.1", "Set scanned ip(default:127.0.0.1)")
	flag.IntVar(&fromPort, "from", 1, "Set the start port(default:1)")
	flag.IntVar(&toPort, "to", 1024, "Set the end port(default:1024)")
}

// 暂且称之为端口扫描工具人
func Worker(ip string, ports, results chan int) {
	// 等待并获取池子中的端口号
	for port := range ports {
		address := fmt.Sprintf("%s:%d", ip, port) // 合成完整的ip地址
		conn, err := net.Dial("tcp", address)     // 尝试与该端口建立tcp连接
		if err != nil {                           // 若连接失败，则表明该端口处于关闭或过滤状态
			results <- 0
			continue // 继续尝试与剩余端口进行连接
		}
		conn.Close()    // 连接成功，无需其他操作，应及时关闭本次连接
		results <- port // 反馈端口号
	}
}

func main() {
	// 这是督工今天从老板那边收到的任务要求
	flag.Parse()

	ports := make(chan int, 100) // 督工提供带缓冲的端口池，准备了100个工作岗位
	results := make(chan int)    // 除了用于接受结果以外，还充当阻塞main goroutine的作用

	opened_ports := []int{} // 督工需要记录所有已开放端口

	// 督工招聘工具人并让它们等待被分配工作
	// 工具人之间的工作是互不干扰的
	for i := 0; i < cap(ports); i++ {
		go Worker(scanIp, ports, results)
	}

	// 督工向端口池中传入端口号，将工作分配给工具人
	go func() {
		for port := fromPort; port <= toPort; port++ {
			ports <- port
		}
	}()

	// 督工等待并接收来自工具人的工作结果
	start := time.Now()
	for port := fromPort; port <= toPort; port++ {
		result := <-results
		if result != 0 { // 同时，记录那些已开放的端口号
			opened_ports = append(opened_ports, result)
		}
	}
	end := time.Since(start)

	close(ports)   // 完成今天的任务后，督工让工具人们下班
	close(results) // 同时，督工自己打卡下班，结束了一天的工作

	sort.Ints(opened_ports) // 晚上，督工按顺序整理今天的任务报告，准备在明天将工作成果向老板你汇报
	for _, port := range opened_ports {
		fmt.Printf("%s:%d is opened.\n", scanIp, port)
	}

	// 工作顺利交差，真是美好的一天
	fmt.Printf("It takes %v\n", end)
}
