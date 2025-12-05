package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main26() {
	s1, s2, s3, s4 := "1", "2", "3", "4"
	merged := strings.Join([]string{s1, s2, s3, s4}, "-")
	fmt.Println(merged)

	builder := strings.Builder{}
	builder.WriteString(s1)
	builder.WriteString(",")
	builder.WriteString(s2)
	builder.WriteString(",")
	builder.WriteString(s3)
	builder.WriteString(",")
	builder.WriteString(s4)
	fmt.Println(builder.String())

	str := strings.Split(merged, "-")
	fmt.Println(str)

	str1 := strconv.Itoa(1)
	fmt.Println(str1)
	// 字符串转整数
	i, err1 := strconv.Atoi(str[1])
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(i)

	// string转int64
	i64, err := strconv.ParseInt(str[2], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i64)

}
