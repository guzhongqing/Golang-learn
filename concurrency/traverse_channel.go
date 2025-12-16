package main

import (
	"fmt"
	"time"
)

var asyncChann = make(chan int, 2)

func ChannelBlock() {
	// 从channel中获取元素
	go func() {
		time.Sleep(time.Second)
		if v, ok := <-asyncChann; ok {
			fmt.Printf("take:%d\n", v)
		}

		// 非阻塞时，等channel获取元素
		time.Sleep(time.Second)

		if v, ok := <-asyncChann; ok {
			fmt.Printf("take:%d\n", v)
		}

	}()
	// 发送元素
	asyncChann <- 100
	fmt.Println("send 100")
	asyncChann <- 200
	fmt.Println("send 200")
	asyncChann <- 300
	fmt.Println("send 300")
	time.Sleep(2 * time.Second)
	close(asyncChann)
	// 关闭channel后，再发送元素，会panic
	// asyncChann <- 400
	// fmt.Println("send 400")

	// 关闭channel后，再从channel中获取元素，会获取channel中剩余的元素
	if v, ok := <-asyncChann; ok {
		fmt.Printf("take:%d\n", v)
	}
	// 关闭channel后，再从channel中获取元素，会返回false
	v, ok := <-asyncChann
	fmt.Printf("take:%d, ok:%v\n", v, ok)

	v1, ok1 := <-asyncChann
	fmt.Printf("take:%d, ok:%v\n", v1, ok1)

}

func TraverseChannel() {
	asyncChann <- 100
	asyncChann <- 200
	// 因为asyncChann定义容量为2，再发送300会死锁
	asyncChann <- 300
	close(asyncChann)
	// 遍历channel中的元素
	// for v := range asyncChann {
	// 	fmt.Printf("take:%d\n", v)
	// }

	for {
		if v, ok := <-asyncChann; ok {
			fmt.Printf("take:%d\n", v)
		} else {
			fmt.Printf("channel closed\n")
			fmt.Printf("take:%d ok:%v\n", v, ok)
			break
		}
	}

}
func main() {
	// ChannelBlock()
	TraverseChannel()

}
