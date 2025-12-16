package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]int
	mu   sync.RWMutex
}

// store 存储键值对
func (cm *ConcurrentMap) store(key string, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.data[key] = value
}

// load 加载键对应的值
func (cm *ConcurrentMap) load(key string) (int, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, ok := cm.data[key]
	return value, ok
}

func NewConcurrentMap(cap int) *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]int, cap),
	}
}

type ConcurrentMapWithGeneric[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func (cm *ConcurrentMapWithGeneric[K, V]) store(key K, value V) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.data[key] = value
}

func (cm *ConcurrentMapWithGeneric[K, V]) load(key K) (V, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, ok := cm.data[key]
	return value, ok
}

func NewConcurrentMapWithGeneric[K comparable, V any](cap int) *ConcurrentMapWithGeneric[K, V] {
	return &ConcurrentMapWithGeneric[K, V]{
		data: make(map[K]V, cap),
	}
}

func TestConcurrentMap() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CollectionSafety", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// cm := NewConcurrentMap(1000)
	// 测试泛型
	// cm := NewConcurrentMapWithGeneric[string, int](1000)
	cm := NewConcurrentMapWithGeneric[int, bool](1000)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i += 2 {
			// cm.store(fmt.Sprintf("key%d", i), i)
			// 测试泛型
			cm.store(i, true)

		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i += 2 {
			// cm.store("key"+strconv.Itoa(i), i)
			// 测试泛型
			cm.store(i, false)

		}
	}()
	// 等待所有 goroutine 完成
	wg.Wait()
	// 打印
	for i := range 1000 {

		// if v, ok := cm.load(fmt.Sprintf("key%d", i)); ok {
		// 	fmt.Printf("%d, value:%v\n", i, v)
		// }
		// 测试泛型
		if v, ok := cm.load(i); ok {
			fmt.Printf("%d, value:%v\n", i, v)
		}
	}
}
func TestConcurrentMapWithGeneric() {
	const P = 10
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done()
			cm := NewConcurrentMapWithGeneric[int, bool](1000)
			for j := 0; j < 1000; j++ {
				cm.store(j, true)
			}
		}()
	}
	// 等待所有 goroutine 完成,没有fatal error
	wg.Wait()
}

// func main() {
// 	// TestConcurrentMap()
// 	TestConcurrentMapWithGeneric()
// }
