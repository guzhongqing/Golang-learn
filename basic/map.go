package main

import "fmt"

func update_map() {
	var m map[string]int     // 声明
	m = make(map[string]int) // 初始化,默认容量为0
	m["a"] = 1
	fmt.Println("m:", m)

	// 这里m重新指向了一个新的map
	m = make(map[string]int, 5) // 初始化,设定容量为5,建议根据元素数量预估容量，减少扩容次数
	m["b"] = 2
	fmt.Println("m:", m)

	m = map[string]int{
		"c": 1,
		"d": 2,
	}

	fmt.Println("m:", m)
	// 定义并赋值
	map1 := map[string]int{
		"e": 3,
		"f": 4,
	}
	fmt.Println("map1:", map1)

	// 添加/更新元素
	map1["g"] = 5
	fmt.Println("map1:", map1)
	map1["g"] = 10
	fmt.Println("map1:", map1)

	// 读取元素
	fmt.Println("map1[e]:", map1["e"])
	fmt.Println("map1[g]:", map1["g"])

	// 删除元素
	delete(map1, "g")
	fmt.Println("map1:", map1)
	// 读取不存在的元素，返回定义map时value的默认值
	fmt.Println("map1[g]:", map1["g"]) // 0是int的默认值

	// 检查元素是否存在,使用逗号 ok 模式comma-ok idiom
	Value, ok := map1["g"]
	if ok {
		fmt.Println("map1[g] value:", Value)
	} else {
		fmt.Println("map1[g] does not exist")
	}
	// 简化检查元素是否存在的代码
	if Value, ok := map1["e"]; ok {
		fmt.Println("map1[e] value:", Value)
	} else {
		fmt.Println("map1[e] does not exist")
	}
	// 获取长度
	fmt.Println("map1 length:", len(map1))

	// 遍历map
	for key, value := range map1 {
		fmt.Println("key:", key, "value:", value)
	}

	// 边遍历边修改
	for key, value := range map1 {
		map1[key] = value * 2
	}
	fmt.Println("map1:", map1)

	// 遍历map,只遍历key
	for key := range map1 {
		fmt.Print(" key:", key)
	}
	fmt.Println()

	// 遍历map,只遍历value,并且修改value,不会原来map的value
	for _, value := range map1 {
		value = value * 2
		fmt.Print(" value:", value)
	}
	fmt.Println()
	fmt.Println(map1)

}

func main28() {
	update_map()

}
