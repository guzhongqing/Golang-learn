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

func function_basic() {
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

func variable_arg(a int, other ...int) {
	sum := a
	for index, element := range other {
		sum += element
		fmt.Println(index, element)
	}
	fmt.Println("sum:", sum)
	if len(other) > 0 {
		fmt.Println("other:", other)
	} else {
		fmt.Println("other is empty")
	}
}

// 递归求和
func sum_recursive(other ...int) int {
	if len(other) == 0 {
		return 0
	}
	sum := other[0]
	// 切片... 表示将切片展开,作为参数传递
	sum += sum_recursive(other[1:]...)
	return sum
}

func main30() {
	// function_basic()
	variable_arg(1, 2, 3, 4, 5)
	variable_arg(1)
	// 递归求和
	fmt.Println(sum_recursive(1, 2, 3, 4, 5))
	fmt.Println(sum_recursive(1))
	fmt.Println(sum_recursive())

}
