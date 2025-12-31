package main

import (
	"fmt"
	"slices"
)

func SliceInit() {

	// 切片声明
	var s []int
	// 切片初始化，空切片
	s = []int{}
	fmt.Println(s)
	// 初始化切片并赋值
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	// len=5, cap=5
	fmt.Printf("切片长度: %d, 切片容量: %d\n", len(s1), cap(s1))

	// 使用make函数创建切片，len=5, cap=10,len的部分会被初始化为默认值0
	s2 := make([]int, 5, 10)
	fmt.Println(s2)
	fmt.Printf("切片长度: %d, 切片容量: %d\n", len(s2), cap(s2))

	// 使用make函数创建切片，len=cap=5
	s3 := make([]int, 5)
	fmt.Println(s3)
	fmt.Printf("切片长度: %d, 切片容量: %d\n", len(s3), cap(s3))

	// 二维切片
	s2d := [][]int{
		{1},
		{2, 3},
	}
	fmt.Println(s2d)
	// 行len=2, cap=2
	fmt.Printf("二维切片长度: %d, 切片容量: %d\n", len(s2d), cap(s2d))
	// 每行的len和cap
	// 第一行len=1, cap=1
	fmt.Printf("二维切片第0个元素len:%d,cap:%d\n", len(s2d[0]), cap(s2d[0]))
	// 第二行len=2, cap=2
	fmt.Printf("二维切片第1个元素len:%d,cap:%d\n", len(s2d[1]), cap(s2d[1]))

}

// 探索切片扩容机制
func Expansion() {
	// 切片扩容
	s := make([]int, 0, 2)
	prevCap := cap(s)
	for i := range 1100 {
		s = append(s, i)
		// fmt.Printf("len=%d, cap=%d", len(s), cap(s))
		if cap(s) != prevCap {
			fmt.Printf("  切片扩容了: %d --> %d,扩容了：%.2f倍\n", prevCap, cap(s), float64(cap(s))/float64(prevCap))
			prevCap = cap(s)
		}

	}
}

// 内存共享
func MemoryShare() {
	s := make([]int, 3, 3)
	s[0], s[1], s[2] = 1, 2, 3
	fmt.Println(s)
	s1 := s
	// 修改s1会影响s
	s1[0] = 100
	fmt.Println("s:", s)
	fmt.Println("s1:", s1)

	// append操作可能会导致底层数组重新分配，从而不再共享内存
	s2 := append(s, 4)
	s2[1] = 200
	fmt.Println("s after append:", s)
	fmt.Println("s1 after append:", s1)
	fmt.Println("s2 after append:", s2)

	// s和s1仍然共享内存
	s1[2] = 300
	fmt.Println("s after modifying s1:", s)
	fmt.Println("s1 after modifying s1:", s1)
	fmt.Println("s2 remains unchanged:", s2)
}

// 切片截断和append
func SliceTruncateAndAppend() {
	// 数组
	var arr [4]int
	fmt.Printf("brr type: %T\n", arr)

	// 通过[a:b]，会返回切片类型,cap=len(arr)-a,容量从a开始计算
	// brr切片的底层数组是arr，与arr共享内存
	brr := arr[1:2]
	fmt.Printf("brr type: %T\n", brr)
	fmt.Printf("brr len:%d,brr cap:%d\n", len(brr), cap(brr))
	// 修改brr会影响arr
	brr[0] = 100
	fmt.Println("arr:", arr)
	fmt.Println("brr:", brr)

	// append只能用于切片类型，不能用于数组类型
	// 扩容之前对brr进行append操作，会修改arr的值
	brr = append(brr, 200)
	fmt.Println("arr after append:", arr)
	fmt.Println("brr after append:", brr)

	// 对arr修改，会影响brr
	arr[2] = 300
	fmt.Println("arr after modifying arr:", arr)
	fmt.Println("brr after modifying arr:", brr)

	// 对arr修改，但是超出brr的cap范围，不会影响brr
	arr[3] = 400
	fmt.Println("arr after modifying arr:", arr)
	fmt.Println("brr after modifying arr:", brr)

	// 对brr进行append操作，会修改arr，导致对应位置的值被覆盖
	brr = append(brr, 500)
	fmt.Println("arr after second append:", arr)
	fmt.Println("brr after second append:", brr)

	// 继续对brr进行append操作，超过cap范围，brr会重新分配内存，不再与arr共享内存
	brr = append(brr, 600)
	fmt.Println("arr after third append:", arr)
	fmt.Println("brr after third append:", brr)

	fmt.Printf("brr len:%d,brr cap:%d\n", len(brr), cap(brr))
}

func SlicesFunction() {
	// 切片的一些函数
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("包含", slices.Contains(arr, 5))
	fmt.Println("最大者", slices.Max(arr))
	fmt.Println("最小者", slices.Min(arr))

	brr := make([]int, 3, 5)
	copy(brr, arr)
	fmt.Println("brr:", brr)

	type User struct {
		Name   string
		Age    int
		Height float64
		Weight float64
	}
	UserSlice := []User{
		{"a", 25, 163, 96},
		{"b", 18, 163, 120},
		{"c", 25, 168, 110},
		{"d", 18, 168, 97},
		{"e", 25, 168, 98},
		{"f", 18, 168, 99},
	}

	for _, v := range UserSlice {
		fmt.Println(v)
	}
	// 按照年龄升序，身高升序，体重降序
	slices.SortFunc(UserSlice, func(a, b User) int {
		if a.Age != b.Age {
			return a.Age - b.Age
		}
		if a.Height != b.Height {
			if a.Height > b.Height {
				return 1
			} else {
				return -1
			}
		}
		if a.Weight != b.Weight {
			if a.Weight > b.Weight {
				return -1
			} else {
				return 1
			}
		}
		return 0
	})
	for _, v := range UserSlice {
		fmt.Println(v)
	}

}

func main25() {
	// SliceInit()
	// Expansion()
	// MemoryShare()
	// SliceTruncateAndAppend()
	SlicesFunction()

}
