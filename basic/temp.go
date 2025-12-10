package main

import "fmt"

type person struct {
	Name string
	Age  int
}
type AInterface interface {
	Add(a, b int) int
	Modify()
}

// 实现AInterface接口
func (p person) Add(a, b int) int {
	return a + b
}

// 没有实现AInterface接口的Sub方法,也不是成员方法
func Sub(a, b int, p person) int {
	return a - b
}

// 该写法和下面的写法完全不等价
// func Modify(p person) {
// 	fmt.Println("modify")
// }

func (p person) Modify() {
	fmt.Println("modify")
}

func (p *person) personAgePlusForPointer() {
	p.Age++
}

func (p person) personAgePlusForStruct() {
	p.Age++
}

func main000() {
	// 结构体的定义,初始化默认值
	var p1 person
	fmt.Println(p1)

	// 调用指针方法，会修改外部结构体的Age字段
	p1.personAgePlusForPointer()
	fmt.Println(p1)

	// 调用结构体方法，不会修改外部结构体的Age字段
	p1.personAgePlusForStruct()
	fmt.Println(p1)

	// 调用接口方法
	var a AInterface = p1
	fmt.Println(a.Add(1, 2))
	// 调用非成员方法
	fmt.Println(Sub(1, 2, p1))

}
