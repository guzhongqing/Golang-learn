package main

import (
	"fmt"
	"math/rand/v2" // 替换为randv2
	"os"
	"path/filepath"
	"time"
)

// 配置项（与原需求一致）
const (
	targetDir    = "./data/biz_log" // 与处理程序目录一致
	fileCount    = 10000            // 生成1.log ~ 10000.log
	minLines     = 1                // 每个文件最少行数
	maxLines     = 100              // 每个文件最多行数
	minNum       = 1                // 数字最小值
	maxNum       = 50               // 数字最大值
	progressStep = 1000             // 每生成1000个文件打印进度
)

// 全局随机数生成器（randv2推荐创建实例而非全局函数）
var rng *rand.Rand

// 初始化randv2生成器（保证每次运行生成不同数据）
func init() {
	// randv2不再使用Seed，而是通过New创建生成器，基于时间+随机数初始化种子
	// NewPCG是randv2推荐的默认生成器（高性能、统计特性好）
	rng = rand.New(rand.NewPCG(
		uint64(time.Now().UnixNano()), // 种子1：当前时间戳
		rand.Uint64(),                 // 种子2：随机64位整数
	))
}

// 生成单个日志文件（randv2实现）
func generateLogFile(filePath string) error {
	// 创建文件（覆盖已有文件，权限0644）
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件%s失败: %v", filePath, err)
	}
	defer file.Close()

	// 随机生成当前文件的行数（1-100）：randv2.IntN与旧版行为一致（返回[0,n)的整数）
	lineCount := rng.IntN(maxLines-minLines+1) + minLines

	// 写入随机数字（每行一个）
	for i := 0; i < lineCount; i++ {
		// 生成1-50的随机整数
		num := rng.IntN(maxNum-minNum+1) + minNum
		// 写入文件（每行一个数字，换行符结尾）
		_, err := fmt.Fprintln(file, num)
		if err != nil {
			return fmt.Errorf("写入文件%s第%d行失败: %v", filePath, i+1, err)
		}
	}

	return nil
}

func main() {
	// 1. 创建目标目录（不存在则创建，存在则忽略）
	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Printf("创建目标目录%s失败: %v\n", targetDir, err)
		return
	}
	fmt.Printf("✅ 目标目录已准备好：%s\n", targetDir)

	// 2. 批量生成10000个日志文件
	fmt.Println("\n开始生成文件（共10000个）...")
	startTime := time.Now()

	for i := 1; i <= fileCount; i++ {
		// 拼接文件路径（如 ./data/biz_log/1.log）
		fileName := fmt.Sprintf("%d.log", i)
		filePath := filepath.Join(targetDir, fileName)

		// 生成当前文件
		err := generateLogFile(filePath)
		if err != nil {
			fmt.Printf("⚠️  生成文件%s失败: %v\n", filePath, err)
			continue
		}

		// 打印进度（每1000个文件一次）
		if i%progressStep == 0 {
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("✅ 已生成%d/%d个文件，耗时%.2f秒\n", i, fileCount, elapsed)
		}
	}

	// 3. 生成完成统计
	totalElapsed := time.Since(startTime).Seconds()
	fmt.Printf("\n🎉 所有文件生成完成！\n")
	fmt.Printf("📂 生成目录：%s\n", targetDir)
	fmt.Printf("📊 生成数量：%d个文件\n", fileCount)
	fmt.Printf("⏱️  总耗时：%.2f秒\n", totalElapsed)
	fmt.Printf("📝 文件规则：每个文件包含%d-%d行，每行是%d-%d的随机整数\n", minLines, maxLines, minNum, maxNum)
}
