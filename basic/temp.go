package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) personAgePlusForPointer() {
	p.Age++
}

func (p person) personAgePlusForStruct() {
	p.Age++
}

func main() {
	// 结构体的定义,初始化默认值
	var p1 person
	fmt.Println(p1)

	// 调用指针方法，会修改外部结构体的Age字段
	p1.personAgePlusForPointer()
	fmt.Println(p1)

	// 调用结构体方法，不会修改外部结构体的Age字段
	p1.personAgePlusForStruct()
	fmt.Println(p1)

}
