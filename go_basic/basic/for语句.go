package main

import "fmt"

func main17() {

	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	// 没有while语句，可以用for来实现
	count := 0
	for count < 10 {
		fmt.Println("Count is:", count)
		count += 2
	}

	// for嵌套
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}

	}

}
