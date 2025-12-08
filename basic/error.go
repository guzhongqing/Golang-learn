package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Name string
	Code int
	Dosc string
}

// Error()方法实现error接口
func (e MyError) Error() string {
	return fmt.Sprintf("[%d]%s", e.Code, e.Dosc)
}

// MyError实现stringer接口的String方法，打印时调用的还是Error()方法，不会调用String()方法
func (e MyError) String() string {
	return "这是String()方法"
}

// 构造函数
func NewMyError(name string, code int, dosc string) error {
	return &MyError{name, code, dosc}
	// 或者使用new关键字
	// err := new(MyError)
	// err.Name = name
	// err.Code = code
	// err.Dosc = dosc
	// return err
}

var (
	ErrNotFound = errors.New("not found error")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		// 使用errors.New()创建错误
		// return 0, ErrNotFound
		// 使用fmt.Errorf()创建错误
		// return 0, fmt.Errorf("divide error a:%d b:%d", a, b)
		// 使用自定义错误
		return 0, NewMyError("divide error", 1001, "divide by zero")
	}
	return a / b, nil
}

func main42() {
	// 测试正常情况
	result, err := divide(10, 5)
	if err != nil {
		// 针对error类型，打印会自动调用Error()方法，这里省略了Error()方法的调用
		fmt.Println("出错:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// 测试除数为0的情况
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}

}
