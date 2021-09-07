/*
文件名批处理器（文件批量重命名）
在指定路径下查找具有特定前缀名称的直属文件
批量剔除前缀名称，以达到重命名文件的作用
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
)

// 平台文件路径分隔符
const PathSeparator = string(os.PathSeparator)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	// 输入要处理的文件所在路径
	fmt.Println("请输入待处理文件的所在路径:")
	path, err := ReadString(inputReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 判断文件路径是否合法
	if err := CheckPath(path); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 将目标路径统一转换为绝对路径，便于表示
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输入待处理文件的名称前缀
	fmt.Println("请输入待处理文件的名称前缀:")
	prefix, err := ReadString(inputReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 根据给定的前缀，匹配目标路径下的文件
	files, err := GetFiles(path, prefix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if len(files) == 0 {
		fmt.Printf("%s路径下暂无以'%s'为前缀的文件!\n", path, prefix)
		return
	}

	fmt.Println("已为您匹配到以下文件:")
	for _, file := range files {
		fmt.Println(joinPath(path, file))
	}

	fmt.Printf("确认是否执行操作[y/n]:")
	verification, err := ReadString(inputReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if verification != "y" {
		fmt.Println("已终止操作!")
		return
	}

	// 处理文件（文件重命名）
	if err := Rename(path, prefix, files); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Done!")
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
func CheckPath(path string) (err error) {
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

// 根据名称前缀，获取指定路径下的所有直属文件名信息
func GetFiles(path, prefix string) (files []string, err error) {
	// 构建指定路径的文件对象
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	// 列出指定路径中所有直属文件信息
	// 根据前缀匹配文件并放置到待处理列表中
	// 如果想获取指定路径下所有文件、子目录及其子文件信息
	// 使用递归或者更高效的path.filepath.Walk()实现
	for _, info := range infos {
		name := info.Name()
		if !info.IsDir() && strings.HasPrefix(name, prefix) {
			files = append(files, name)
		}
	}

	return
}

// 重命名文件
func Rename(path, prefix string, files []string) (err error) {
	// 获取原文件名
	for _, filename := range files {
		// 根据前缀切分原文件名并生成新文件名
		new_filename := filename[len(prefix):]

		old_path := joinPath(path, filename)
		new_path := joinPath(path, new_filename)

		if err = os.Rename(old_path, new_path); err != nil {
			return
		}
	}
	// done!
	return
}

// 文件路径拼接
func joinPath(paths ...string) string {
	return strings.Join(paths, PathSeparator)
}
