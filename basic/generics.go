package main

import (
	"cmp"
	"fmt"
	"reflect"
)

type Age1 int32
type Int32Float64 interface {
	~int32 | float64
}

func getBiggerForInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBiggerForFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 使用泛型
// func getBiggerForAny[T int32 | float64](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 使用接口定义的泛型
//
//	func getBiggerForAny[T Int32Float64](a, b T) T {
//		if a > b {
//			return a
//		}
//		return b
//	}
//
// 使用go自带的cmp包泛型约束
func getBiggerForAny[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 结构体使用泛型
// 定义一个结构体，包含一个泛型字段
type Apple[T cmp.Ordered] struct {
	Value T
}

// 定义结构体方法时，在结构体后面添加[T cmp.Ordered]，表示该方法是一个泛型方法
func (Apple[T]) getBigger(a, b T) T {
	if a > b {
		return a
	}
	return b
}

type GetUserRequest struct{}
type GetBookRequest struct{}

// 不使用泛型
// func httpRPC(request any ) {

// 使用泛型会在编译时检查类型是否符合约束
func httpRPC[T GetUserRequest | GetBookRequest](request T) {
	tf := reflect.TypeOf(request)
	switch tf.Name() {
	case "GetUserRequest":
		fmt.Println("get user")
	case "GetBookRequest":
		fmt.Println("get book")
	}

}

func main50() {
	biggerForFloat64 := getBiggerForFloat64(1, 2)
	fmt.Println(biggerForFloat64)
	// 调用泛型函数，如果泛型是int，可以省略，因为无类型数字字面量可以自动推断为默认整数int
	biggerForInt := getBiggerForAny[int32](3, 4)
	fmt.Println(biggerForInt)

	// 调用泛型函数，传入Age1类型，因为Age1是int32的别名，所以可以传入int32类型的参数
	biggerForAge1 := getBiggerForAny[Age1](5, 6)
	fmt.Println(biggerForAge1)

	apple1 := Apple[int32]{Value: 1}
	fmt.Println(apple1.getBigger(7, 8))

	apple2 := Apple[float32]{Value: 1.0}
	fmt.Println(apple2.getBigger(1.0, 2.0))

	// 调用httpRPC函数，传入GetUserRequest类型的参数
	httpRPC(GetUserRequest{})
	httpRPC(GetBookRequest{})
	// 编译错误，int类型不符合约束
	// httpRPC(4)

}
