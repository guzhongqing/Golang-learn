package main

import "fmt"

// 类型别名，需要强制类型转换
type Age int

// 类型别名, 完全等价，不需要显式类型转化，但是语法上也可以使用
type Tall = int

// Age 实现了 Transporter 接口
func (c Age) move(src string, dest string) (int, error) {
	fmt.Println("Age move")
	return 0, nil
}

func (c Age) whistle(n int) int {
	fmt.Println("Age whistle")
	return 0
}
func (c *Age) test() {
	fmt.Println("测试给结构体指针绑定方法")

}

func main() {
	var a int
	var b Age
	var c Tall
	b.test()
	(&b).test()


	c = a + 10
	b = Age(a) + 10 // 字面量不需要强制类型转换
	fmt.Println(a, b, c)

	// 存储的数据类型（成员变量）一致，但是行为（成员方法）不一样
	type shipType Ship
	var s shipType
	s.Name = "泰坦尼克"
	s.Price = 2000000
	s.Tonage = 100000
	fmt.Println(s)

	// 显式把shipType转换为Ship类型，shipType的行为和Ship类型的行为是一致的
	ship := Ship(s)
	// 调用 Ship 类型的方法
	ship.move("北京", "上海")
	ship.whistle(2)

	// Age 实现了 Transporter 接口
	transport(b, "北京", "上海")

}
