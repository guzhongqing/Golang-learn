package main

import "fmt"

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

func main() {
	// SliceInit()
	Expansion()
}
