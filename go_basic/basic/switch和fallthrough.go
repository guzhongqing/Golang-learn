package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main19() {

	color := "red"
	switch color {
	case "yellow":
		fmt.Println("The color is yellow")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("The color is", color)

	}

	switch color {
	case "yellow":
		fmt.Println("The color is yellow")
	case "red":
		fmt.Println("The color is red")
	default:
		fmt.Println("The color is", color)
	}

	switch color {
	case "yellow", "red":
		fmt.Println("The color is yellow or red")
	default:
		fmt.Println("The color is", color)
	}

	// switch判断函数
	switch add(1, 2) {
	case 1:
		fmt.Println("Result is 1")
	case 2:
		fmt.Println("Result is 2")
	case 3:
		fmt.Println("Result is 3")
	default:
		fmt.Println("Result is unknown")
	}

	// switch fallthrough 强制执行下一个case 或default
	switch color {
	case "yellow":
		fmt.Println("The color is yellow")
		fallthrough
	case "red":
		fmt.Println("The color is red")
		fallthrough
	default:
		fmt.Println("The default color is bule")
	}

}
