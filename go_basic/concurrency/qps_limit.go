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
	// 模拟接口调用耗时，qps限流之后网络请求、数据库操作，这些中间件的qps就会被限流到，但是该接口的qps是不受影响的
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
