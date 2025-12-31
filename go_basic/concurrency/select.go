package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用select监听多个通道
func ListenMultiWay() {

	ch1 := make(chan int, 1000)
	ch2 := make(chan byte, 1000)

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch1 <- rand.Int()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch2 <- byte(rand.Intn(256))
		}
	}()

	// 这里AB标签只绑定在for循环上，下面last的select语句不会阻塞
AB:
	for {
		time.Sleep(time.Second)
		select {
		case v1 := <-ch1:
			fmt.Println("ch1:", v1)
		case v2 := <-ch2:
			fmt.Println("ch2:", v2)
			if v2 < 40 {
				break AB
			}
		// 当ch1和ch2都没有数据时，执行default,不会阻塞，没有default会阻塞
		default:
			fmt.Println("default")
		}
	}

	select {
	case v1 := <-ch1:
		fmt.Println("at last ch1:", v1)
	default:
	}

}

func SelectBlock() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("子协程还在工作")
		}
	}()

	// 阻塞main协程，等待子协程执行完毕，但是会一直阻塞
	select {}

}

// func main() {
// 	// ListenMultiWay()
// 	SelectBlock()
// }
