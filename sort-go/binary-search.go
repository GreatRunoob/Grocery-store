/*
二分查找算法，针对拥有n个已排好序的元素集合
*/

package sort

type sortType int

func BinarySearch(list []sortType, target sortType) int {
	var low int = 0
	var high int = len(list) - 1

	for {
		if low <= high {
			mid := (low + high) / 2
			if target == list[mid] {
				return mid
			} else if target < list[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			break
		}
	}

	return -1
}
