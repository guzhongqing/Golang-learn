package main

import (
	"fmt"
	"sync"
)

type Student struct {
	Name string
	Age  int
}

var arr = [10]int{}

var m = sync.Map{}

func SafetyStruct() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CollectionSafety", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	var student Student

	// 只修改Name
	go func() {
		defer wg.Done()
		student.Name = "张三"

	}()
	// 只修改Age
	go func() {
		defer wg.Done()
		student.Age = 18

	}()
	// 等待所有 goroutine 完成
	wg.Wait()
	// 打印
	fmt.Println(student)
}

func SafetyArr() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CollectionSafety", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	// 偶数位写0
	go func() {
		defer wg.Done()
		for i := 0; i < len(arr); i += 2 {
			arr[i] = 0
		}
	}()
	// 偶数位写1
	go func() {
		defer wg.Done()
		for i := 1; i < len(arr); i += 2 {
			arr[i] = 1
		}
	}()
	// 等待所有 goroutine 完成
	wg.Wait()
	// 打印
	fmt.Println(arr)
}

// 可能会出现fatal error
func UnSafetyMap() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CollectionSafety", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	var mp = make(map[int]bool, 10000)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i += 2 {
			mp[i] = true
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 10000; i += 2 {
			mp[i] = false
		}
	}()
	// 等待所有 goroutine 完成
	wg.Wait()
	// 打印
	fmt.Println(mp)
}

func SafetyMap() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CollectionSafety", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i += 2 {
			m.Store(i, true)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i += 2 {
			m.Store(i, false)
		}
	}()
	// 等待所有 goroutine 完成
	wg.Wait()
	// 打印
	for i := range 1000 {
		if v, ok := m.Load(i); ok {
			fmt.Printf("key:%d, value:%v\n", i, v)
		}
	}

}

func CollectionSafety() {
	// SafetyStruct()
	// SafetyArr()
	// UnSafetyMap()
	SafetyMap()

}

// func main() {
// 	CollectionSafety()

// }
