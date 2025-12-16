package main

import (
	"fmt"
	"time"
)

func Broadcast() {
	ch := make(chan struct{})
	const P = 3

	for i := 0; i < P; i++ {
		go func() {
			<-ch
			fmt.Printf("take:%d\n", i)
		}()
	}
	fmt.Println("send signal")
	close(ch)
	// 等待所有goroutine执行完毕
	time.Sleep(3 * time.Second)
}

// 等同于waitGroup
func CountDownLatch() {
	ch := make(chan struct{})
	const P = 3

	for i := 0; i < P; i++ {
		go func() {
			// 模拟工作,这里i是变量所以样类型强转
			time.Sleep(time.Duration(i) * time.Second)
			// 编译报错
			// time.Sleep(i * time.Second)

			fmt.Printf("完成工作了:%d\n", i)

			ch <- struct{}{}
		}()
	}
	// 等待所有goroutine执行完毕
	for i := 0; i < P; i++ {
		<-ch
	}
	// 所有goroutine执行完毕
	fmt.Println("所有goroutine执行完毕")
}
func main() {
	// Broadcast()
	CountDownLatch()
}
