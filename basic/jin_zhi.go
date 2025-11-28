package main

import "fmt"

/*
多行

注释
这种很少见
*/

// excel最后一列是XFD，求excel有多少列
func main4() {
	fmt.Printf("A=%d Z=%d\n", 'A', 'Z')
	var base int = 'Z' - 'A' + 1
	fmt.Println("进制", base)
	var total int
	total += 'D' - 'A' + 1
	total += ('F' - 'A' + 1) * base
	total += ('X' - 'A' + 1) * base * base
	fmt.Println("总列数", total)

}
