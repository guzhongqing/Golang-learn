package main

import "fmt"

func main9() {
	// 定义单个常量，可以不使用，只会告警
	const PI = 3.14

	// 定义多个常量，枚举
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)

	fmt.Println("StatusOK:", StatusOK)

	// 使用 常量生成器 iota 定义枚举常量
	const (
		a = iota // 0
		b        // 1
		c = 7    // 7
		d        // 7
		e = iota // 4 跳过 c 和 d
		f        // 5
	)
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e, "f:", f)

	// 使用位运算定义权限常量
	const (
		PermRead    = 1 << iota // 1 << 0 which is 0001 输出1
		PermWrite               // 1 << 1 which is 0010 输出2
		PermExecute             // 1 << 2 which is 0100 输出4
	)
	fmt.Println("PermRead:", PermRead, "PermWrite:", PermWrite, "PermExecute:", PermExecute)

	// 常量重复使用上一个值或者表达式
	const (
		const1 = 20          // 20
		const2               // 20
		const3               // 20
		const4 = const1 + 10 // 30
		const5               // 30
		const6               // 30
	)
	fmt.Println("const1:", const1, "const2:", const2, "const3:", const3, "const4:", const4, "const5:", const5, "const6:", const6)

	// 位移使用iota
	const (
		_  = iota             // 忽略第一个值
		KB = 1 << (10 * iota) // 1 << (10*1) 输出1024
		MB                    // 1 << (10*2) 输出1048576
		GB                    // 1 << (10*3) 输出1073741824
		TB                    // 1 << (10*4) 输出1099511627776
		PB                    // 1 << (10*5) 输出1125899906842624
	)
	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB, "TB:", TB, "PB:", PB)

	// 多个常量在一行使用iota
	const (
		ss, tt = iota + 1, iota + 2 // ss=1, tt=2
		uu, vv                      // uu=2, vv=3
		ww, xx                      // ww=3, xx=4
	)
	fmt.Println("ss:", ss, "tt:", tt, "uu:", uu, "vv:", vv, "ww:", ww, "xx:", xx)

}
