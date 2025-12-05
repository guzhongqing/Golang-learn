package main

import (
	"fmt"
	"net/http"
)

type ConnectionPool struct {
	Servers []string
	// 定义一个函数 类型的成员变量 LoadBalancer
	LoadBalancer func(a, b int, c string, d bool) (int, bool)
}
type FT func(a, b int, c string, d bool) (int, bool)

type ConnectionPool1 struct {
	Servers []string
	// 定义一个函数 类型的成员变量 LoadBalancer
	LoadBalancer FT
}

// 给FT类型添加成员方法
// FT 实现了 Transporter 接口
func (ft FT) move(src string, dest string) (int, error) {
	fmt.Println("FT move")
	return 0, nil
}

func (ft FT) whistle(n int) int {
	fmt.Println("FT whistle")
	return 0
}

func main38() {
	// 定义函数类型变量 f1, f2
	var f1, f2 func(a, b int, c string, d bool) (int, bool)

	// 给 f1 赋值一个匿名函数
	f1 = func(a, b int, c string, d bool) (int, bool) {
		return a + b, c == "abcd" && d
	}
	// 使用_占位，忽略f2
	_ = f2

	// 使用 f1 调用匿名函数
	result1, ok1 := f1(1, 1, "abcd", true)
	fmt.Println(result1, ok1) // 输出: 3 true

	// 定义一个 ConnectionPool 结构体

	// 初始化 ConnectionPool 实例
	cp := ConnectionPool{
		Servers: []string{"192.168.1.1", "192.168.1.2"},
		// 给 LoadBalancer 赋值一个匿名函数
		LoadBalancer: f1,
	}
	result2, ok2 := cp.LoadBalancer(1, 2, "abcd", true)
	fmt.Println(result2, ok2) // 输出: 3 true

	// 初始化 ConnectionPool1 实例
	cp1 := ConnectionPool1{
		Servers: []string{"192.168.1.1", "192.168.1.2"},
		// 给 LoadBalancer 赋值一个匿名函数
		LoadBalancer: FT(f1),
	}

	// 使用 cp1.LoadBalancer 调用匿名函数
	result3, ok3 := cp1.LoadBalancer(2, 3, "abcd", true)
	fmt.Println(result3, ok3) // 输出: 5 true

	// 调用 FT 类型的方法
	// cp1.LoadBalancer.move("北京", "上海")
	// cp1.LoadBalancer.whistle(2)

	transport(FT(f1), "北京", "上海")

	// http包调用方法和上面transport方法类似
	h := func(w http.ResponseWriter, r *http.Request) {}
	http.ListenAndServe(":8080", http.HandlerFunc(h))
	//func ListenAndServe(addr string, handler Handler) 第二个参数是接口，要出传入这个接口的实现，一个类型实现了Handler接口的所有方法，就可以作为第二个参数传入
	// 函数h没有实现http.Handler接口的ServeHTTP方法，所以不能作为第二个参数传入
	// 但是http.HandlerFunc类型实现了ServeHTTP方法，把func(ResponseWriter, *Request)强转为http.HandlerFunc类型
	http.ListenAndServe(":8080", http.HandlerFunc(h))

}
