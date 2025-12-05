package main

import "fmt"

type Selector interface {
	Select([]string) int
}

// 面向接口编程
type ConnectionPool3 struct {
	Servers []string
	// 定义一个接口类型的成员变量 LoadBalancer
	LoadBalancer Selector
	// 这里可以直接修改LoadBalancer为InterLeave类型
	LoadBalancer1 InterLeave
}

func f1([]string) int {
	fmt.Println("调用了f1函数")
	return 1
}

func f2([]string) int {
	fmt.Println("调用了f2函数")
	return 2
}

type RoundRobin struct{}

// RoundRobin 结构体实现 Selector 接口的 Select 方法，因为RoundRobin对象没有成员变量，直接省略了变量名
func (RoundRobin) Select(s []string) int {
	fmt.Println("调用了RoundRobin.Select方法")
	// 通过调用f1函数来实现Select方法
	return f1(s)
}

// 更换接口
type InterLeave struct{}

func (InterLeave) Select(s []string) int {
	fmt.Println("调用了InterLeave.Select方法")
	// 通过调用f2函数来实现Select方法
	return f2(s)
}

// 面向函数编程
type ConnectionPool4 struct {
	Servers []string
	// 定义一个函数 类型的成员变量 LoadBalancer
	LoadBalancer func([]string) int
}

func main39() {
	// 初始化 ConnectionPool3 实例
	cp3 := ConnectionPool3{
		Servers: []string{"192.168.1.1", "192.168.1.2"},
		// 给 LoadBalancer 接口，赋值一个 RoundRobin 实例
		LoadBalancer: RoundRobin{},
		// 给 LoadBalancer1 接口，赋值一个 InterLeave 实例
		LoadBalancer1: InterLeave{},
	}
	RoundRobinResult := cp3.LoadBalancer.Select(cp3.Servers)
	fmt.Println(RoundRobinResult) // 输出: 1
	InterLeaveResult := cp3.LoadBalancer1.Select(cp3.Servers)
	fmt.Println(InterLeaveResult) // 输出: 2

	// 初始化 ConnectionPool4 实例
	cp4 := ConnectionPool4{
		Servers: []string{"192.168.1.1", "192.168.1.2"},
		// 给 LoadBalancer 赋值一个函数
		LoadBalancer: f1,
		// 给 LoadBalancer1 赋值一个函数
		// LoadBalancer1: f2,
	}
	result4 := cp4.LoadBalancer(cp4.Servers)
	fmt.Println(result4) // 输出: 1

}
