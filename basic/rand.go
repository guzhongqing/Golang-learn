package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var (
	// 定义一个数字字母表，用于生成随机字符串
	letterbytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// 定义一个数字字母中文表，用于生成随机字符串
	letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ你好")
)

func GenerateRandomString(max int) string {
	// 使用切片存储
	// randStrSlice := make([]rune, max)

	// strings.Builder 存储
	randStrBuilder := strings.Builder{}

	// 生成长度为 max 的随机字符串
	for range max {
		// 定义随机索引
		randIndex := rand.IntN(len(letterRunes))
		// randStrSlice = append(randStrSlice, letterRunes[randIndex])
		randStrBuilder.WriteRune(letterRunes[randIndex])
	}
	// return string(randStrSlice)
	return randStrBuilder.String()
}

func randBasic() {
	// 创建一个新的 PCG 实例(PCG是一个现代、高效的伪随机数生成算法)
	source := rand.NewPCG(123, 456)
	// 生成 5 个随机数
	for range 5 {
		// 使用 Seed 方法设置随机数种子相同时，生成的随机数序列是相同的
		source.Seed(124, 456)
		randNumber := rand.New(source)
		fmt.Println(randNumber.IntN(100))
	}

	// 直接使用 IntN 方法生成随机数
	fmt.Println(rand.IntN(100))

}

func main53() {
	// 调用 randBasic 函数
	// randBasic()

	// 生成随机字符串
	fmt.Println(GenerateRandomString(10))
}
