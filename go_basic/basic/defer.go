package main

import (
	"fmt"
	"time"
)

func deferBasic() {
	fmt.Println("a")
	defer fmt.Println("1")

	fmt.Println("b")
	defer fmt.Println("2")

	fmt.Println("c")
	defer fmt.Println("3")

}

// 定义函数返回值时，声明返回变量为retrnValue，称为带命名返回值的函数
func deferExeTime() (retrnValue int) {
	retrnValue = 9
	// 匿名函数在defer中执行
	defer func() {
		fmt.Printf("first retrnValue=%d\n", retrnValue)
	}()

	defer func(i int) {
		fmt.Printf("second i=%d\n", i)
	}(retrnValue)

	defer fmt.Printf("thire retrnValue=%d\n", retrnValue)
	return 5
}

// 使用defer优化函数调用时间计算
func timeOfWork(arg int) int {
	begin := time.Now()
	defer func() {
		fmt.Printf("timeOfWork cost %v\n", time.Since(begin).Seconds())
	}()
	if arg > 10 {
		time.Sleep(2 * time.Second)
		// 使用defer优化重复的代码
		// fmt.Printf("timeOfWork cost %v\n", time.Since(begin).Seconds())
		return 100
	} else {
		time.Sleep(3 * time.Second)
		// fmt.Printf("timeOfWork cost %v\n", time.Since(begin).Seconds())
		return 200
	}

}
func main32() {
	// deferBasic()
	// deferExeTime()
	// num := timeOfWork(5)
	num := timeOfWork(20)
	fmt.Printf("num=%d\n", num)

}
