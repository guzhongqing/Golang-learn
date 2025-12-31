package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// 定义空结构体
type ETS struct{}

// 所有空结构体指向同一个地址
func allEmptyStructIsSame() {
	var a ETS
	var b ETS
	var c struct{}
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	// 打印空结构体的大小
	fmt.Printf("Size of ETS: a=%d, b=%d, c=%d\n", unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
	fmt.Printf("Size of ETS: a=%d, b=%d, c=%d\n", reflect.TypeOf(a).Size(), reflect.TypeOf(b).Size(), reflect.TypeOf(c).Size())
	// 查看正常结构体
	type NormalStruct struct {
		a int
		b string
	}
	normal := NormalStruct{
		a: 1,
		b: "hello",
	}
	fmt.Printf("%p\n", &normal)

	fmt.Printf("Size of NormalStruct: %d 字节\n", unsafe.Sizeof(normal))

}

func scenariosOfEmptyStruct() {
	// 使用空结构体实现set结构
	set := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
	}
	if _, ok := set[1]; ok {
		fmt.Println("1 is in set")
	}
	if _, ok := set[3]; !ok {
		fmt.Println("3 is not in set")
	}

	blocker := make(chan struct{}, 0)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("dome")
		// 往channel写数据，通知主goroutine继续执行
		blocker <- struct{}{}

	}()
	// 主goroutine阻塞等待channel有数据
	<-blocker
	fmt.Println("main continue")

}

func main41() {
	// allEmptyStructIsSame()
	scenariosOfEmptyStruct()
}
