package main

import (
	"fmt"
)

type Transporter interface {
	// 只定义方法签名，不定义变量
	move(src string, dest string) (int, error)
	whistle(int) int
}

type Streamer interface {
	Transporter // 通过匿名成员实现接口"继承"
	// displacementwei 计算运输距离
	displacement() int
}

type Car struct {
	Name  string
	Price int
}

// Car 实现了 Transporter 接口
// 只要结构体拥有接口里声明的所有方法，就称该结构体"实现了接口"
func (c Car) move(src string, dest string) (int, error) {
	fmt.Printf("Car %s move from %s to %s need cost %d\n", c.Name, src, dest, c.Price)
	return c.Price, nil
}

func (c Car) whistle(n int) int {
	for i := range n {
		fmt.Printf("car滴：%d\n", i)
	}
	return n
}

// 可以有额外的方法
func (c Car) GetName() string {
	return c.Name
}

//

type Ship struct {
	Name  string
	Price int
	// 载重吨位
	Tonage int
}

// Ship 实现了 Streamer和 Transporter 接口
func (s Ship) move(src string, dest string) (int, error) {
	fmt.Printf("Ship %s move from %s to %s need cost %d\n", s.Name, src, dest, s.Price)
	return s.Price, nil
}

func (s Ship) whistle(n int) int {
	for i := range n {
		fmt.Printf("ship滴：%d\n", i)
	}
	return n
}
func (s Ship) displacement() int {
	return s.Tonage
}

// 可以有额外的方法
// GetName 是导出方法（首字母大写），编译器会默认认为这个方法「可能被其他包导入后调用」，
// 因此即便你在当前 main 包的 interfaceBasic()、main() 中都没用到它，编译器也不会判定为 “未使用”，自然不会提示。
// getName 是未导出方法（首字母小写），它只能在当前 main 包内被调用，如果在当前包未使用，编译器会提示 "getName is unused"
func (s Ship) getName() string {
	return s.Name
}

// 接口作为函数形参，调用时可以传入任意实现了该接口的结构体实例
// PS：接口也可以作为结构体的成员变量（策略模式）
func transport(t Transporter, src string, dest string) error {
	if _, err := t.move(src, dest); err != nil {
		return err
	} else {
		t.whistle(3)
		return nil
	}

}

func seaTransport(s Streamer, src string, dest string) error {
	fmt.Printf("Shi displacement is %d\n", s.displacement())
	if _, err := s.move(src, dest); err != nil {
		return err
	} else {
		s.whistle(3)
		return nil
	}
}

func interfaceBasic() {
	car := Car{
		Name:  "奔驰",
		Price: 1000000,
	}
	ship := Ship{
		Name:   "泰坦尼克",
		Price:  2000000,
		Tonage: 100000,
	}

	// 调用 transport 函数时，传入不同的结构体实例，实现了多态
	transport(car, "北京", "上海")
	transport(ship, "北京", "上海")

	// 定义接口, 实现了 Transporter 接口的结构体实例都可以赋值给该接口变量
	var transporter Transporter
	transporter = car
	transporter.whistle(2)

	transporter = ship
	transporter.whistle(2)

	fmt.Println()

	seaTransport(ship, "北京", "上海")

}

func main36() {
	interfaceBasic()

}
