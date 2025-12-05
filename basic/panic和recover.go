package main

import "fmt"

func deferPanic() {
	// 注册的defer会在panic触发后执行
	defer fmt.Println("a")
	arr := make([]int, 0, 3)
	arr = append(arr, 1, 2, 3)
	fmt.Println(arr)
	n := 0
	defer func() {
		// 正常函数执行完，才会执行defer函数的内容
		_ = arr[3]
		// 不会执行
		_ = 1 / n
		// 下面的defer不会执行
		defer fmt.Println("b")
	}()
	// 这个已经注册
	defer fmt.Println("外部", arr)
}

func recoverPanic() {
	// 捕获panic，防止程序崩溃，执行到该函数时，有panic，才会捕获，如果执行过后，才出现panic，不会捕获到
	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			fmt.Println("recoverPanic", panicInfo)
		}
	}()
	// 触发panic，会导致defer执行
	panic("触发panic")
}
func main33() {
	fmt.Println("deferPanic")
	// 直接调用panic函数，会导致程序崩溃
	// panic(1)
	// deferPanic()
	recoverPanic()
	fmt.Println("继续执行")

}
