package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func New(text string) error {
	return &MyError{text}
}

var (
	ErrNotFound = errors.New("not found error")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrNotFound
	}
	return a / b, nil
}

func main() {
	// 测试正常情况
	result, err := divide(10, 2)
	if err != nil {
		// 针对error类型，打印会自动调用Error()方法，这里省略了Error()方法的调用
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// 测试除数为0的情况
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}

}
