package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

func add2(a int, b int) int {
	return a + b
}

func return1() (int, int) {
	return 1, 2
}

func return2(a int) (int, int) {
	return a + 1, a + 2
}

// 指针作为函数参数
func arg1(a, b *int) {
	*a = *a + *b
	*b = 888
}

func main21() {
	fmt.Println(add1(1, 1))
	fmt.Println(add2(1, 1))
	fmt.Println(return1())
	a, b := return2(1)
	fmt.Println(a, b)
	x, y := 3, 5
	// 指针作为函数参数,会改变实参的值
	arg1(&x, &y)
	fmt.Println(x, y)

}
