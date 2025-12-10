package main

import (
	"fmt"
)

// 定义函数类型
type Adder func() int

// addd 是一个函数，它返回一个函数，这个函数的参数是 int，返回值也是 int
func addd() Adder {
	a := 10
	// 打印地址
	fmt.Printf("a 的地址是 %p\n", &a)
	return func() int {
		a++
		fmt.Printf("a 的地址是 %p\n", &a)

		return a
	}
}

type FuncNoParamNoReturn func()

func funcList() {
	var funcSlice []func()
	for i := 0; i < 3; i++ {
		// 打印地址
		fmt.Printf("i 的地址是 %p\n", &i)
		funcSlice = append(funcSlice, func() {
			// 在返回时，使用i，此时返回函数会持有i
			fmt.Printf("i 的地址是 %p\n", &i)
			println(i)

		})
	}
	// 调用 funcSlice 中的函数
	for _, f := range funcSlice {
		f()
	}
}

func closureBasic() {
	// 调用 addd 函数，返回一个 Adder 类型的函数
	af := addd()
	// 调用 adder 函数，返回 11
	fmt.Println(af())
	// 调用 adder 函数，返回 12
	fmt.Println(af())
	bf := addd()
	fmt.Println(bf())

}

func main48() {
	// closureBasic()

	funcList()
}
