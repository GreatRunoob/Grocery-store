/*
有关选择排序的基本原理简单描述一下：
每趟排序从剩余的序列中选择极值，插入到当前序列头部
并在下一轮排序中，将该极值排除在外。
*/
package main

import "fmt"

type sortType int

// 查找并返回序列中最小元素索引值，传入查找序列
func findSmallest(arr []sortType) int {
	// 假设序列首元素为本轮排序的极值，并记录其索引值。以查找最小值元素为例
	smallest_index := 0
	smallest := arr[smallest_index]

	// 尝试寻找本轮排序其他可能的最小值及其索引值
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest_index = i
			smallest = arr[i]
		}
	}

	return smallest_index
}

// 使用选择排序，对给定序列进行升序排序并返回排序号后的新序列
func selectionSort(arr []sortType) []sortType {
	count := len(arr)
	result := []sortType{}

	for i := 0; i < count; i++ {
		smallest_index := findSmallest(arr)
		result = append(result, arr[smallest_index])

		// 本轮排序，剔除当前序列中的极值元素
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}

	return result
}

func main() {
	// 测试样例
	arr := []sortType{3, 1, 9, 6, 0, 5, 7, 2, 8, 4}
	arr = selectionSort(arr)
	fmt.Println(arr)
}
