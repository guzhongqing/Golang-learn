package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var n int32 = 0

func inc1() {
	// 非原子操作
	n++

}

func inc2() {
	// 原子操作
	atomic.AddInt32(&n, 1)

}
func Atomic() {
	const N = 1000
	var wg = sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			// defer fmt.Printf("协程 %d 结束\n ", i)

			// 非原子操作
			// inc1()
			// 原子操作
			// inc2()
			// 加锁操作
			inc3()
		}()
	}
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())
	// 等待所有协程结束
	wg.Wait()
	fmt.Printf("n 的值：%d\n", n)
	fmt.Printf("当前运行的 goroutine 数量：%d\n", runtime.NumGoroutine())

}

var lock sync.Mutex

func inc3() {
	lock.Lock()   // 加互斥锁
	n++           // 临界区，任意时刻只有一个 goroutine 可以执行
	lock.Unlock() //释放互斥锁
}

// 定义读写锁
var mu sync.RWMutex

// 读锁可重入
func ReentranceRLock(n int) {
	mu.RLock()         // 加读锁
	defer mu.RUnlock() // 释放读锁
	fmt.Println(n)
	if n > 0 {
		ReentranceRLock(n - 1)
	}
	time.Sleep(time.Second)
}

// 写锁不可重入
func ReentranceLock(n int) {
	mu.Lock()         // 加写锁
	defer mu.Unlock() // 释放写锁
	fmt.Println(n)
	if n > 0 {
		ReentranceRLock(n - 1)
	}
	time.Sleep(time.Second)
}

// 读锁的排他性
func RLockExclusion() {
	mu.RLock() // 加读锁
	defer func() {
		mu.RUnlock() // 释放读锁
		fmt.Println("主协程 goroutine 释放读锁1")
	}()
	go func() {
		mu.RLock() // 加读锁

		defer func() {
			mu.RUnlock() // 释放读锁
			fmt.Println("子 goroutine 释放读锁")
		}()
		fmt.Println("子 goroutine 加读锁")
	}()

	// // 加读锁不会阻塞
	// mu.RLock() // 加读锁
	// defer func() {
	// 	mu.RUnlock() // 释放读锁
	// 	fmt.Println("当前 goroutine 释放读锁2")
	// }()

	go func() {
		mu.Lock()         // 加写锁
		defer mu.Unlock() // 释放写锁
		fmt.Println("子 goroutine 加写锁")
	}()

	// 主协程加写锁会死锁,因为目前所有协程（当前只有一个主协程）都被阻塞
	// mu.Lock()         // 加写锁
	// defer mu.Unlock() // 释放写锁
	// fmt.Println("当前 goroutine 加写锁")

	time.Sleep(5 * time.Second)

}

// 写锁的排他性
func WLockExclusion() {
	mu.Lock()         // 加写锁
	defer mu.Unlock() // 释放写锁

	go func() {
		mu.RLock()         // 加读锁
		defer mu.RUnlock() // 释放读锁
		// 长时间会导致子协程获取到锁时，主协程已经结束，子协程强制结束
		// time.Sleep(2 * time.Second)
		fmt.Println("子 goroutine 加读锁")
	}()

	// go func() {
	// 	mu.Lock()         // 加写锁
	// 	defer mu.Unlock() // 释放写锁
	// 	fmt.Println("子 goroutine 加写锁")
	// }()

	// 死锁
	// mu.RLock() // 加读锁
	// fmt.Println("主协程 goroutine 加读锁")
	// mu.RUnlock() // 释放读锁
	// fmt.Println("当前 goroutine 释放读锁2")

	// 死锁
	// mu.Lock()         // 加写锁
	// defer mu.Unlock() // 释放写锁
	// fmt.Println("当前 goroutine 加写锁")
	// time.Sleep(5 * time.Second)

}

func main() {
	// Atomic()
	// ReentranceRLock(3)
	// ReentranceLock(3)
	// RLockExclusion()
	WLockExclusion()

}
