package main

import (
	"fmt"
	"runtime"
	"time"
)

// 定义限流结构体
type GoroutineLimiter struct {
	limit int
	ch    chan struct{}
}

// 构造函数
func NewGoroutineLimiter(limit int) *GoroutineLimiter {
	return &GoroutineLimiter{
		limit: limit,
		ch:    make(chan struct{}, limit),
	}
}

// 限流方法
func (g *GoroutineLimiter) Run(f func()) {
	g.ch <- struct{}{}

	go func() {
		f()
		<-g.ch
	}()
}

func GoroutineLimit() {
	tk := time.NewTicker(1 * time.Second)
	defer tk.Stop()

	go func() {
		for {
			<-tk.C
			// 每一秒打印当前goroutine数量
			fmt.Printf("当前goroutine数量: %d\n", runtime.NumGoroutine())
		}
	}()

	// 模拟工作协程
	work := func() {
		// 模拟工作
		time.Sleep(1 * time.Second)
	}
	const P = 100
	// 正常启动P个工作协程
	// for i := 0; i < P; i++ {
	// 	go work()
	// }

	// 限流启动P个工作协程
	// 创建一个限流结构体
	g := NewGoroutineLimiter(10)
	for i := 0; i < P; i++ {
		g.Run(work)
	}

	// 等待所有工作协程完成
	time.Sleep(10 * time.Second)
}

// func main() {
// 	GoroutineLimit()
// }
