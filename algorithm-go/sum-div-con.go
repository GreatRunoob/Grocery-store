/*
分而治之，利用递归实现切片求和
*/
package main

import "fmt"

func Sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + Sum(a[1:])
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	fmt.Printf("Sum of array is: %d\n", Sum(a))
}
