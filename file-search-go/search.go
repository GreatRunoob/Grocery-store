/*
文件/目录查找工具
使用指定的关键词，模糊搜索目标路径下所有可能的匹配项
1.开启适当数量的协程有助于提升搜索速度
2.合适的搜索工作分配算法有助于提高平均搜索效率
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const PathSeparator = string(os.PathSeparator) // 平台文件路径分隔符
const maxWorkerCount int = 32                  // 最大同时进行路径搜索的协程数

var searchRequest = make(chan string) // 请求另外启用路径搜索协程
var foundMatch = make(chan string)    // 存在可能的匹配项
var workingDone = make(chan bool)     // 协程本轮搜索工作完成记录

var matches int = 0     // 共计匹配数
var workerCount int = 0 // 当前搜索协程开启数量

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	// 绝对路径和相对路径都可以
	fmt.Println("请输入目标文件路径:")
	path, err := ReadString(inputReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 检查目标文件路径是否合法
	if err := PathCheck(path); err != nil {
		fmt.Println("Error", err)
		return
	}

	// 将目标路径统一转换为绝对路径
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 模糊查找关键词
	fmt.Println("请输入待搜索关键词:")
	substr, err := ReadString(inputReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 搜索计时
	start := time.Now()

	// 启动第一个搜索协程
	workerCount++
	go Search(path, substr, true) // 从目标根路径开始，按照关键词搜索匹配项

	// 统计相关数据、分配工作并等待所有所搜协程工作结束(wait for workers)
	for {
		select {
		case path := <-searchRequest: // 接收到搜索分工请求
			workerCount++                 // 将当前正在进行搜索工作的协程数量+1
			go Search(path, substr, true) // 启动单独的协程分担搜索工作
		case <-workingDone: // 接收到协程工作结束信号
			workerCount--         // 讲当前正在进行搜索工作的协程数量-1
			if workerCount == 0 { // 当前不存在进行搜索工作的协程
				goto end // 结束等待，搜索工作结束
			}
		case path := <-foundMatch: // 接收到可能的匹配项
			fmt.Println(path)
			matches++ // 匹配记录+1
		}
	}

end:
	close(searchRequest)
	close(foundMatch)
	close(workingDone)

	if matches != 0 {
		fmt.Printf("已为您匹配到以上内容，耗时:%v\n", time.Since(start))
		fmt.Printf("目标匹配数:%d\n", matches)
	} else {
		fmt.Printf("暂未找到可能的匹配项！耗时:%v\n", time.Since(start))
	}
}

// 读取标准输入
func ReadString(inputReader *bufio.Reader) (input string, err error) {
	input, err = inputReader.ReadString('\n')
	if err != nil {
		return
	}

	input = input[:len(input)-1] // 剔除换行符
	return
}

// 路径检查
func PathCheck(path string) (err error) {
	// 判断给定路径是否存在
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return
	}

	// 且是否为文件目录
	if !stat.IsDir() {
		return errors.New(path + " is not a directory!")
	}

	return
}

// 模糊搜索
func Search(path, substr string, goroutine bool) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, substr) { // 匹配是否包含关键词
			if file.IsDir() {
				foundMatch <- pathJoin(path, filename, "")
				// e.g. foundMatch <- "/root/dir/"
			} else {
				foundMatch <- pathJoin(path, filename)
				// e.g. foundMatch <- "/root/file"
			}
		}

		if file.IsDir() { //未匹配到关键词，但本身是子目录
			if workerCount < maxWorkerCount { // 检查当前是否允许提交分工请求
				searchRequest <- pathJoin(path, filename) // 允许，提交请求
			} else {
				Search(pathJoin(path, filename), substr, false) // 不允许，让当前协程递归进行搜索
			}
		}
	}

	// 直到当前路径下所有子文件全部进行过匹配
	// 若当前协程匹配的是普通文件，则无需继续递归查询

	if goroutine { // 协程本次搜索工作结束之前，传递结束信号
		workingDone <- true
	}
	return
}

// 文件路径拼接
func pathJoin(paths ...string) string {
	return strings.Join(paths, PathSeparator)
}
