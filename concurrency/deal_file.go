package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 定义各功能协程数
const (
	WALK_FILE_GOROUTINE_COUNT = 1
	READ_FILE_GOROUTINE_COUNT = 15
	PROCESS_GOROUTINE_COUNT   = 5
)

var (
	sum      int64
	fileList = make(chan string, 100)
	// lineBuffer通道，用于存储读取的行内容
	lineBuffer = make(chan string, 1000)

	walkWg    sync.WaitGroup
	readWg    sync.WaitGroup
	processWg sync.WaitGroup
)

// 递归遍历dir目录，将所有文件路径发送到fileList通道
func walkDir(dir string) {
	defer walkWg.Done()
	// 遍历目录,dir为绝对路径，path也是绝对路径，dir为相对路径，path为相对路径
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("需要处理的文件路径：", path)
		// 只处理普通文件(如 .txt、.go、.jpg 等可读写的常规文件)
		// 非普通文件(如目录、符号链接、管道、设备文件、套接字等)，则跳过
		if !info.Mode().IsRegular() {
			return nil
		}
		// 普通文件，发送到fileList通道
		fileList <- path
		return nil
	})
}

// 从fileList通道读取文件路径，读取文件内容发送到lineBuffer通道
func readFile() {
	defer readWg.Done()

	// 一个readFile协程会一直从fileList通道读取文件路径，直到fileList通道关闭
	// 通道关闭且通道为空时，for循环结束
	for path := range fileList {
		// 读取文件内容
		fin, err := os.Open(path)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return
		}
		defer fin.Close()
		// 使用bufio读取行
		// 创建bufio.Reader对象
		reader := bufio.NewReader(fin)
		// 循环读取行，直到文件结束，
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// 如果最后一行没有换行符，也将其发送到lineBuffer通道
					if line != "" {
						lineBuffer <- strings.TrimSpace(line)
					}
					// 到达文件末尾，正常退出循环
					break
				}
				// 其他错误，打印错误信息
				fmt.Println("读取文件失败:", err)
				return
			}
			lineBuffer <- strings.TrimSpace(line)

		}
	}

}

func processLine() {
	defer processWg.Done()
	// 遍历lineBuffer通道，处理行内容
	for line := range lineBuffer {
		// line转换为整数
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("转换整数失败:", err)

		} else {
			// 原子累加整数
			atomic.AddInt64(&sum, int64(num))
		}
	}
}

func DealMassFile(dir string) {
	// 开启一个协程周期查看通道数量
	go func() {
		tk := time.NewTicker(time.Second)
		defer tk.Stop()
		for {
			<-tk.C
			fmt.Printf("fileList通道数量堆积了%d文件未处理\n", len(fileList))
			fmt.Printf("lineBuffer通道数量堆积了%d行未处理\n", len(lineBuffer))
		}
	}()

	walkWg.Add(WALK_FILE_GOROUTINE_COUNT)
	readWg.Add(READ_FILE_GOROUTINE_COUNT)
	processWg.Add(PROCESS_GOROUTINE_COUNT)
	// 启动WALK_FILE_GOROUTINE_COUNT个walkDir协程
	for i := 0; i < WALK_FILE_GOROUTINE_COUNT; i++ {
		go walkDir(dir)
	}
	// 启动READ_FILE_GOROUTINE_COUNT个readFile协程
	for i := 0; i < READ_FILE_GOROUTINE_COUNT; i++ {
		go readFile()
	}
	// 启动PROCESS_GOROUTINE_COUNT个processLine协程
	for i := 0; i < PROCESS_GOROUTINE_COUNT; i++ {
		go processLine()
	}
	// 等待所有walkDir协程完成
	walkWg.Wait()
	// 关闭fileList通道，通知readFile协程停止读取
	close(fileList)
	// 等待所有readFile协程完成
	readWg.Wait()
	// 关闭lineBuffer通道，通知processLine协程停止处理
	close(lineBuffer)
	// 等待所有processLine协程完成
	processWg.Wait()

	// 打印累加结果
	fmt.Println("累加结果:", sum)
}

// func main() {
// 	// 待处理的目录
// 	// dir := "./data/biz_log"
// 	dir := `D:\code\vscode\go_project\go_basic\concurrency\data\biz_log`
// 	// 处理目录下的所有文件
// 	DealMassFile(dir)
// }
