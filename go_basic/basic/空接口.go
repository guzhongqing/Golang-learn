package main

import "fmt"

type EmptyInterface interface{}

// func SumI(args ...interface{}) int {
// func SumI(args ...any) int {
func SumI(args ...EmptyInterface) int {

	sum := 0
	for _, ele := range args {
		switch eleType := ele.(type) {
		case int:
			sum += eleType //此时已经是确定类型 eleType 是 int 类型
		case float32:
			sum += int(eleType) // 此时已经是确定类型 eleType 是 float32 类型
		default:
			fmt.Printf("不支持%T类型，值为%v\n", eleType, ele)
		}

	}

	return sum
}
func main40() {
	sum := SumI(1, 2.0, float32(3.14), "3", true)
	fmt.Println(sum)
	fmt.Println()

	// 定义空接口类型的map
	mp := make(map[any]interface{}, 10)
	mp["a"] = 1
	mp["b"] = 2.0
	mp["c"] = float32(3.14)
	mp["d"] = "3"
	mp["e"] = true
	mp[9] = 18

	for key, value := range mp {
		fmt.Printf("key type: %[1]T %[1]v, value type: %[2]T %[2]v\n", key, value)
	}
	fmt.Println()

	// 定义空接口类型的slice
	sl := make([]EmptyInterface, 0, 10)
	sl = append(sl, 10, 2.0, float32(3.14), "1", true)
	fmt.Printf("sum of slice: %d\n", SumI(sl...))

	// 类型断言
	var i any
	if v, ok := i.(int); ok {
		fmt.Printf("i is int type, value is %d\n", v)
	} else {
		fmt.Println("i is not int type")
	}

	var j any = 100
	if v, ok := j.(int); ok {
		fmt.Printf("j is int type, value is %d\n", v)
	} else {
		fmt.Println("j is not int type")
	}

}
