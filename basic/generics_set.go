package main

import (
	"errors"
	"fmt"
)

// 使用内置comparable接口约束泛型类型T
type Set[T comparable] map[T]struct{}

// 定义构造函数
func NewSet[T comparable](cap int) Set[T] {
	return make(Set[T], cap)
}

// 定义长度获取方法
func (s Set[T]) Len() int {
	return len(s)
}

// 定义添加元素方法
func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}

// 定义删除元素方法
func (s Set[T]) Del(k T) {
	// 先判断元素是否存在，再删除
	if _, ok := s[k]; !ok {
		err := errors.New("元素不存在")
		fmt.Println(err)
		return
	}
	delete(s, k)
}

// 定义判断元素是否存在方法
func (s Set[T]) Exist(k T) bool {
	_, ok := s[k]
	return ok
}

// 实现String方法，用于打印集合元素
func (s Set[T]) String() string {
	keys := make([]T, 0, s.Len())
	for k := range s {
		keys = append(keys, k)
	}
	return fmt.Sprintf("%v", keys)
}

// 定义遍历元素方法
func (s Set[T]) Range(f func(ele T)) {
	for k := range s {
		f(k)
	}
}

func main51() {
	set := NewSet[int](10)
	fmt.Println(set.Len())
	set.Add(1)
	fmt.Println(set.Len())

	// 判断元素是否存在
	fmt.Println(set.Exist(1))
	fmt.Println(set.Exist(2))

	set.Del(2)
	fmt.Println(set.Len())
	// 删除已存在元素
	set.Del(1)
	fmt.Println(set.Len())
	// 添加元素
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(2)

	// 打印
	fmt.Println(set)

	// 遍历元素
	set.Range(func(ele int) {
		// 定义回调函数
		fmt.Printf("%d\n", ele)
	})

}
