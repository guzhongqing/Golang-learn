package main

// 按行并发读取一个文件
// 1.一次性读取文件内容到切片，再按照数组并发处理
// 2.针对读取的行数，进行记录，计算行数时要加锁
// 3.修改成流式读取文件内容，通过通道发送数据，避免内存占用过大，同时支持并发处理
import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func read_file(filePath string, lineChan chan<- string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 流式读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChan <- scanner.Text()
	}

	close(lineChan) // 关闭通道，表示数据发送完毕

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}
	return nil
}

// 按行并发读取一个文件
func file_concurrency() {
	filePath := "./file_concurrency.txt"
	maxCap := 2
	lineChan := make(chan string, 100) // 缓冲通道，减少阻塞
	goroutinePool := make(chan struct{}, maxCap)
	lineNumber := 0

	var wg sync.WaitGroup
	var mutex sync.Mutex

	// 启动读取协程
	go func() {
		if err := read_file(filePath, lineChan); err != nil {
			panic(err)
		}
	}()

	// 处理数据
	idx := 0
	for value := range lineChan {
		// 增加正在执行的协程的数量和协程池限制的计数
		wg.Add(1)
		goroutinePool <- struct{}{}

		go func(idx int, value string) {
			// 结束要减掉正在执行的数量，并且减掉协程池里面的数量
			defer func() {
				wg.Done()
				<-goroutinePool
				// 记录文件处理的行数
				mutex.Lock()
				lineNumber++
				mutex.Unlock()
			}()

			// 处理文件流
			fmt.Printf("索引:%d,值:%s\n", idx, value)

		}(idx, value)
		idx++
	}

	wg.Wait()
	fmt.Printf("共处理 %d 行数据\n", lineNumber)

}
