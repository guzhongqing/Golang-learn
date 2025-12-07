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

	p1.personAgePlusForPointer()
	fmt.Println(p1)

	p1.personAgePlusForStruct()
	fmt.Println(p1)

}
