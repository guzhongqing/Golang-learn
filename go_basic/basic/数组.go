package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func array1D() {
	// 数组的定义,初始化默认值
	var arr1 [5]int
	fmt.Println(arr1)

	// 数组的定义,初始化赋值省略类型，根据值自动推导
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 数组的定义,初始化赋值省略长度，根据值自动推导
	arr3 := [...]int{4, 5, 6, 7}
	fmt.Println(arr3)

	// 数组的定义,指定索引赋值，未赋值的元素使用默认值
	arr4 := [5]int{0: 10, 3: 30}
	fmt.Println(arr4)

	// 数组定义，给前两个赋值
	var arr5 [5]int = [5]int{1, 2}
	fmt.Println(arr5)

	// 结构体数组
	arr6 := [3]Person{
		{"Alice", 30},
		{"Bob", 25},
	}
	fmt.Println(arr6)

	// 匿名结构体数组
	arr7 := [...]struct {
		name string
		age  int
	}{
		{"Charlie", 28},
	}
	fmt.Println(arr7)

	// 通过索引访问数组元素
	fmt.Println(arr6[0])
	// 访问最后一个元素
	fmt.Println(arr3[len(arr3)-1])

	fmt.Printf("数组的地址：%p\n", &arr3)
	fmt.Printf("数组第一个元素的地址：%p\n", &arr3[0])
	fmt.Printf("数组第二个元素的地址：%p\n", &arr3[1])

	// 遍历数组
	for i, ele := range arr2 {
		fmt.Println("index:", i, "element:", ele)
	}

	// 只遍历元素
	for _, ele := range arr3 {
		fmt.Println("element:", ele)
	}
	// for 循环遍历，传统方式，建议使用 range
	for i := 0; i < len(arr4); i++ {
		fmt.Println("index:", i, "element:", arr4[i])
	}

	// 数组的长度是确定不变，capacity 和 length 相等
	fmt.Println("数组arr4的长度:", len(arr4), "容量:", cap(arr4))

}

func array2D() {
	// 二维数组
	// 定义一个5行3列的二维数组,
	array1 := [5][3]int{
		{1},
		{2, 3},
	}
	fmt.Println(array1)

	// 定义一个二维数组，第1维使用[...]省略长度，由编译器推导，第2维不能用...
	array2 := [...][3]int{
		{1},
		{2, 3},
	}
	fmt.Println(array2)

	// 访问二维数组元素
	fmt.Println(array1[1][1]) // 输出3

	// 遍历二维数组
	for rowIndex, rowArray := range array2 {
		fmt.Println("rowIndex:", rowIndex, "rowArray:", rowArray)
		for colIndex, ele := range rowArray {
			fmt.Println("row:", rowIndex, "col:", colIndex, "element:", ele)
		}
	}

	// 对于二维数组，len和cap返回的是第1维的长度，即行数
	fmt.Println("数组arr4的长度:", len(array2), "容量:", cap(array2))

}

func updateArray1(arr [3]int) {
	fmt.Printf("函数内数组地址:%p\n", &arr)
	arr[0] = 100
	fmt.Println("函数内修改后的数组:", arr)
}
func updateArray2(arr *[3]int) {
	fmt.Printf("函数内数组地址:%p\n", &(*arr))
	arr[1] = 100 // 通过指针原修改数组内容
	fmt.Println("函数内修改后的数组:", *arr)
}

// 传指针的数组
func updateArray3(arr [3]*int) {
	fmt.Printf("函数内数组地址:%p\n", &arr)
	*arr[2] = 100 // 通过指针原修改数组内容
	fmt.Println("函数内修改后的数组:", arr)

}

// 数组传参给函数
func funcArray() {
	arr := [...]int{1, 2, 3}
	fmt.Printf("函数外数组地址:%p\n", &arr)
	updateArray1(arr)
	fmt.Println("updateArray1函数外数组值:", arr)

	// 传数组指针
	updateArray2(&arr)
	fmt.Println("updateArray2函数外数组值:", arr)

	// 传指针的数组
	PointArray := [3]*int{&arr[0], &arr[1], &arr[2]}
	fmt.Printf("函数前:\nPointArray[0]:%d\nPointArray[1]:%d\nPointArray[2]:%d\n", *PointArray[0], *PointArray[1], *PointArray[2])
	updateArray3(PointArray)
	fmt.Println("updateArray3函数外数组值:", arr)
	fmt.Printf("函数后:\nPointArray[0]:%d\nPointArray[1]:%d\nPointArray[2]:%d\n", *PointArray[0], *PointArray[1], *PointArray[2])

}

// for range 取的是拷贝值
func forRangeArray() {
	arr := [...]int{10, 20, 30}
	for index, element := range arr {
		arr[index] += 100
		element += 8
		fmt.Printf("index:%d, element:%d, arr[%d]:%d\n", index, element, index, arr[index])
	}
	fmt.Println("最终数组值:", arr)
}

func main22() {
	// array1D()
	// array2D()
	// funcArray()
	forRangeArray()

}
