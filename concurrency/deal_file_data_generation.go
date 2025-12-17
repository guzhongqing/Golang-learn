package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

// 配置项
const (
	targetDir    = "./data/biz_log" // 目标目录
	fileCount    = 20000            // 总文件数(1.log~10000.log)
	minLines     = 1                // 单文件最少行数
	maxLines     = 30               // 单文件最多行数
	minNum       = 1                // 数字最小值
	maxNum       = 50               // 数字最大值
	progressStep = 1000             // 每生成1000个文件打印进度
	workerCount  = 4                // 优化：HDD设4~6，SSD设8~10（而非固定10）
)

// 生成单个日志文件（优化：批量写入+bufio减少syscall）
func generateLogFile(filePath string, rng *rand.Rand) error {
	// 创建文件（O_CREATE|O_WRONLY|O_TRUNC 等价于os.Create，但显式控制）
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("创建失败: %w", err)
	}
	defer file.Close()

	// 优化：用bufio.Writer缓冲写入，减少syscall次数（批量写内存，再刷到磁盘）
	writer := bufio.NewWriterSize(file, 4096) // 4KB缓冲（匹配磁盘块大小）
	defer writer.Flush()                      // 最后刷入磁盘

	// 随机行数(1-100)
	lineCount := rng.IntN(maxLines-minLines+1) + minLines
	var content []byte // 预分配内存，减少字符串拼接

	// 批量拼接内容（内存操作，无IO）
	for i := 0; i < lineCount; i++ {
		num := rng.IntN(maxNum-minNum+1) + minNum
		content = append(content, []byte(fmt.Sprintf("%d\n", num))...)
	}

	// 一次写入缓冲（仅1次syscall）
	if _, err := writer.Write(content); err != nil {
		return fmt.Errorf("写入失败: %w", err)
	}

	return nil
}

// 工作协程（优化：减少日志打印频率，降低锁竞争）
func worker(taskChan <-chan int, wg *sync.WaitGroup, progress *atomic.Int64, workerID int) {
	defer wg.Done()

	// 每个协程独立rng（避免竞争）
	rng := rand.New(rand.NewPCG(
		uint64(time.Now().UnixNano())+uint64(workerID)*100, // 唯一种子
		rand.Uint64()+uint64(workerID)*100,
	))

	for fileNum := range taskChan {
		filePath := filepath.Join(targetDir, fmt.Sprintf("%d.log", fileNum))

		if err := generateLogFile(filePath, rng); err != nil {
			// 优化：仅每100个错误打印一次，避免刷屏+锁竞争
			if fileNum%100 == 0 {
				fmt.Printf("[协程%d] ⚠️  %s: %v\n", workerID, filePath, err)
			}
			continue
		}

		// 原子更新进度
		completed := progress.Add(1)
		// 优化：仅让一个协程打印进度（减少stdout锁竞争）
		if completed%progressStep == 0 && workerID == 1 {
			fmt.Printf("[进度] ✅ 已生成%d/%d个文件\n", completed, fileCount)
		}
	}
}

func runGenerateFile() {
	// 1. 创建目标目录（优化：提前检查权限）
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 目标目录：%s\n", targetDir)

	// 2. 初始化任务通道（优化：缓冲设为workerCount*100，匹配协程处理能力）
	taskChan := make(chan int, workerCount*100)
	var progress atomic.Int64
	var wg sync.WaitGroup

	// 3. 启动协程（优化：根据磁盘类型调整workerCount）
	fmt.Printf("\n启动%d个协程生成%d个文件...\n", workerCount, fileCount)
	startTime := time.Now()
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(taskChan, &wg, &progress, i+1)
	}

	// 4. 分发任务（优化：无缓冲分发，避免内存占用）
	go func() {
		for i := 1; i <= fileCount; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	// 5. 等待完成
	wg.Wait()

	// 6. 统计
	totalElapsed := time.Since(startTime).Seconds()
	fmt.Printf("\n🎉 生成完成！\n")
	fmt.Printf("📊 数量：%d个文件 | ⚡️ 协程数：%d\n", fileCount, workerCount)
	fmt.Printf("⏱️  总耗时：%.2f秒 | 📝 平均速度：%.0f文件/秒\n",
		totalElapsed, float64(fileCount)/totalElapsed)
	fmt.Printf("📝 规则：每个文件%d-%d行，每行%d-%d的整数\n", minLines, maxLines, minNum, maxNum)
}

// func main() {
// 	// 优化：提前预热文件系统缓存（可选）
// 	_, _ = os.Stat(targetDir)
// 	runGenerateFile()
// }
