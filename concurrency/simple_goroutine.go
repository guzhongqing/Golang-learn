package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Add(a, b int) {
	fmt.Printf("Add 函数被调用，参数 a=%d, b=%d\n", a, b)
}

func SimpleGoroutine() {
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())

	fmt.Printf("当前逻辑核心数：%d\n", runtime.NumCPU())
	// 修改最大逻辑核心数
	maxProcs := runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	// 查看修改后的最大逻辑核心数
	fmt.Printf("修改后的最大逻辑核心数maxProcs：%d\n", maxProcs)
	fmt.Printf("修改后的最大逻辑核心数runtime.GOMAXPROCS(0)：%d\n", runtime.GOMAXPROCS(0))

	// 执行一个 goroutine
	go Add(1, 2)
	go Add(1, 2)

	// 调用匿名函数
	go func(a, b int) {
		fmt.Printf("匿名函数被调用，参数 a=%d, b=%d\n", a, b)
	}(1, 2)
	// 定义并赋值函数变量
	var addFunc = func(a, b int) {
		fmt.Printf("函数变量被调用，参数 a=%d, b=%d\n", a, b)
	}
	// 执行函数变量
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)
	go addFunc(1, 2)

	// 查看当前运行的 goroutine 数量
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
}

func grandson() {
	// 函数结束后，计数器减1
	defer wg.Done()

	fmt.Println("grandson start")
	// 查看当前运行的 goroutine 数量
	fmt.Printf("grandson当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)
	fmt.Printf("grandson当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	fmt.Println("grandson end")
}

func child() {
	// 函数结束后，计数器减1
	defer wg.Done()

	fmt.Println("child start")
	// 查看当前运行的 goroutine 数量
	fmt.Printf("child当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	go grandson()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("child当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	fmt.Println("child end")
}

func SubRoutine() {
	// 定义计数器
	wg.Add(2)
	go child()
	// time.Sleep(2 * time.Second)
	// 等待计数器为0
	wg.Wait()
	fmt.Println("SubRoutine end")

}

func WaitGroup() {
	const N = 10
	var wg = sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(a, b int) {
			defer wg.Done()
			defer fmt.Printf("协程 %d 结束\n ", i)
			time.Sleep(5 * time.Second)
			_ = a + b
		}(i, i+1)
	}
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	// 等待所有协程结束
	wg.Wait()
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())

}

// func main() {
// 	// SimpleGoroutine()
// 	// SubRoutine()
// 	WaitGroup()

// }
