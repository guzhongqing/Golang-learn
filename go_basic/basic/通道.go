package main

import (
	"fmt"
)

func channelBasic() {
	// 声明channel
	var channel chan int
	if channel == nil {
		fmt.Println("channel is nil len:", len(channel), "cap:", cap(channel))
	}
	// 初始化channel
	channel = make(chan int, 8)
	if channel != nil {
		fmt.Println("channel is not nil len:", len(channel), "cap:", cap(channel))
	}
	// 向channel中写入(send函数)数据
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	channel <- 5
	fmt.Println("channel len:", len(channel), "cap:", cap(channel))

	// 从channel中取走(recv函数)数据
	fmt.Println(<-channel)

	// 关闭channel后，不能再向channel中写入数据
	close(channel)
	// 需要先关闭channel，才能遍历channel中的数据，否则会导致deadlock
	for ele := range channel {
		fmt.Println(ele)
	}

	ch := make(chan int, 8)
	send(ch)
	recv(ch)
	sendAndRecv(ch)
}

// 定义只写函数 chan<- 是只写类型
func send(ch chan<- int) {
	ch <- 100
}

// 定义只读函数 <-chan 是只读类型
func recv(ch <-chan int) {
	v := <-ch
	fmt.Println(v)
}

// 定义读写函数 chan 是可读可写类型
func sendAndRecv(ch chan int) {
	ch <- 200
	fmt.Println(<-ch)
}

func main34() {
	channelBasic()

}
