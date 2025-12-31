package main

import "fmt"

func main11() {
	fmt.Println("Go 语言支持继承，但不是通过类继承，而是通过匿名结构体的嵌套来实现。")
	type User struct {
		Name string
		Age  int
	}

	// 匿名结构体
	type AnonymousUser struct {
		Name string
		Age  int
	}

	type Vedio struct {
		Length        int
		Name          string
		Author        User // 结构体嵌套
		AnonymousUser      // 匿名成员

	}
	u := User{
		Name: "Alice",
		Age:  18,
	}
	v := Vedio{
		Length: 120,
		Name:   "Go 语言教程",
		Author: u,
	}
	fmt.Println("视频名称:", v.Name)
	fmt.Println("视频时长:", v.Length)
	fmt.Println("视频作者:", v.Author)
	fmt.Println("作者姓名:", v.Author.Name)
	fmt.Println("作者年龄:", v.Author.Age)

	// 通过匿名结构体成员访问字段
	v.Age = 25
	fmt.Println(v.Age)                // 看成Video从User“继承”了Age字段
	fmt.Println(v.AnonymousUser.Name) // 访问“父类”的Name字段
	fmt.Println(v.Name)               // 访问自己的Name字段

}
