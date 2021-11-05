package sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, elem := range arr[1:] {
		if elem <= privot {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}

	return append(
		QuickSort(left),
		append(
			[]int{pivot}, QuickSort(right)...,
		)...,
	)
}
