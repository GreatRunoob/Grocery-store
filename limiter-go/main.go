package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建token请求限速为3 tokens/s，桶大小为5 tokens的限速器
	limiter := NewTokenLimiter(3, 5)
	for {
		// 模拟并发请求，4 tokens/s
		sameTime := 4
		for i := 1; i <= sameTime; i++ {
			go func(num int) {
				if !limiter.Allow() {
					fmt.Printf("Forbid [%d]\n", num)
				} else {
					fmt.Printf("Allow [%d]\n", num)
				}
			}(i)
		}
		time.Sleep(time.Second)
		fmt.Println("=================")
	}
}
