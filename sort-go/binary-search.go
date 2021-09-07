/*
二分查找算法，针对拥有n个已排好序的元素集合
*/

package main

type sortType int

// 二分查找，返回序列中目标元素索引值，查找失败返回-1
func BinarySearch(list []sortType, target sortType) int {
	var low int = 0
	var high int = len(list) - 1

	for {
		if low <= high {
			mid := (low + high) / 2
			if target == list[mid] {
				return mid
			} else if target < list[mid] {
				high = mid - 1 // 缩小查找范围的上限
			} else {
				low = mid + 1 // 提高查找范围的下限
			}
		} else {
			break // 查找失败
		}
	}

	return -1
}
