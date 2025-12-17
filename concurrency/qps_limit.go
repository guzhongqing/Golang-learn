package main

import (
	"fmt"
	"sync"
	"time"
)

// 限制QPS为100
var qpsCh = make(chan struct{}, 10)

// 模拟处理函数，每次处理耗时3秒
func handle() {
	// 调用接口时，先入队
	qpsCh <- struct{}{}

	// 调用结束，出队
	defer func() {
		<-qpsCh
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("接口调用结束")

}

func limitQPS() {
	const P = 25
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done()
			handle()
		}()
	}
	// 等待所有协程完成
	wg.Wait()

}

// func main() {
// 	limitQPS()

// }
