/*
使用go实现一个极为简单的令牌桶限流器实验模型
用于实现对服务器后端突发请求的限流保护机制
只是简单的模拟，具体实现与理论模型还有很大差别

基本原理：
令牌桶具有固定大小，限流器以恒定的速率向桶中投放token令牌
当令牌桶中的token已满时，暂时不投放token
用户从令牌桶中获取token，如果有剩余的token获取到之后放行
如果没有剩余token，则需要令牌桶中被投放新的token才能放行
*/
package main

import (
	"sync"
	"time"
)

type TokenLimiter struct {
	rate   float64 // 令牌传入限制速率，时间单位s
	tokens float64 // 令牌桶中，当前令牌的数量
	burst  int     // 令牌桶大小

	mutex sync.Mutex // 确保并发安全
	last  time.Time  // 记录最近一次消耗限流令牌的时间
}

func NewTokenLimiter(rate float64, burst int) *TokenLimiter {
	return &TokenLimiter{
		rate:  rate,
		burst: burst,
	}
}

func (this *TokenLimiter) AllowN(now time.Time, requests int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 计算最近一次请求以来，理论上需要补充的token数量
	delta := now.Sub(this.last).Seconds() * this.rate
	this.tokens += delta

	// 检查补充完token之后是否超出桶的容量，多余的token被丢弃
	if this.tokens > float64(this.burst) {
		this.tokens = float64(this.burst)
	}

	// 检查当前token请求量是否超出桶内剩余token数量
	if this.tokens < float64(requests) {
		return false
	}

	// 满足限速的token请求，更新桶内剩余token数量以及最近一次token请求时间
	this.tokens -= float64(requests)
	this.last = now
	return true
}

func (this *TokenLimiter) Allow() bool {
	return this.AllowN(time.Now(), 1)
}
