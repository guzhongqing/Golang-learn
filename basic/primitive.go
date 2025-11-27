package main

import "fmt"

func main2() {
	var MyName int
	fmt.Println(MyName)
	var age int = 30
	fmt.Println(age)
	var b = age // 自动类型推导
	_ = b
	c := b + 1 //第一次出现 短声明 a:= 等同于 var a =
	age = c    //第二次出现不能有声明符号
	fmt.Println(age)

	number := 1_000_000 // 数字字面量中的下划线
	fmt.Println(number)

	m := 13.13 // 浮点型, := 默认是float64
	fmt.Println(m)

	var isOk bool = true // 布尔型, 默认值false
	fmt.Println(isOk)

	var (
		d uint8 // 别名type
		e int8
		f float32
		g float64
	)
	_, _, _, _ = d, e, f, g

}
