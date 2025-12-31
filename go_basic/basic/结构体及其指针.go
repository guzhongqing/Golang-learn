package main

import "fmt"

type User struct {
	Id            int
	Age           int
	score         float64
	Name, Address string
}

// 成员方法
func (mi User) hello() {
	fmt.Println("Hello, my name is", mi.Name)
}

func main13() {
	var u User
	// 结命名赋值法,没有赋值的字段会被初始化为默认值
	u = User{
		Id:   1,
		Age:  18,
		Name: "Alice",
	}
	fmt.Println(u)
	// 按顺序赋值法
	u = User{2, 25, 88.8, "Bob", "New York"}
	fmt.Println(u)
	// 后续赋值
	u.Address = "Los Angeles"
	fmt.Println(u)

	// 调用成员方法
	u.hello()

	// 匿名结构体，通过var和:=定义，不需要事先定义结构体类型
	var student struct {
		Name  string
		Grade int
	}
	fmt.Println(student) // 默认值
	student.Name = "David"
	student.Grade = 10
	fmt.Println(student)

	person := struct {
		Name string
		Age  int
	}{
		Name: "Charlie",
		Age:  30,
	}
	fmt.Println(person)

	u1 := &u             // & 获取结构体变量u的地址，u1是一个指向User类型的指针
	fmt.Println(u1)      // 输出结构体变量u的地址
	fmt.Println(u1.Name) // 通过指针访问结构体的字段

	user := new(User) // new分配内存，创建空的结构体，返回指向User类型的指针
	fmt.Println(user)

	// 结构体嵌套
	type Residence struct {
		Province string
		City     string
	}
	type Person struct {
		Name      string
		Age       int
		Residence Residence
		// 结构体其中可以包含指向自身类型的指针，实现类似链表，二叉树等数据结构，默认值为nil
		Father *Person
		// 匿名结构体
		Contact struct {
			Phone string
			Email string
		}
	}
	p := Person{
		Name: "Eve",
		Age:  28,
		Residence: Residence{
			Province: "California",
			City:     "San Francisco",
		},
	}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

}
