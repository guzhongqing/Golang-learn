package main

import (
	"strconv"
)

var table = []int{1, 2, 3, 4, 5, 6, 7}

func GetHandler(index int) int {
	defer func() {
		// 捕获panic
		recover()
	}()
	return table[index]

}

func SetHandler(index int, div string) {
	defer func() {
		// 捕获panic
		recover()
	}()
	i, _ := strconv.Atoi(div)
	table[index] = 10 / i

}

// func main() {
// 	// 当前协程只捕获当前goroutine的panic
// 	// defer func() {
// 	// 	// 捕获panic
// 	// 	recover()
// 	// }()
// 	go GetHandler(10)
// 	go SetHandler(1, "a")
// 	time.Sleep(time.Second)
// }
