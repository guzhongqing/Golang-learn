package main

func isValid(s string) bool {

	// 首先判断字符串长度，如果是奇数，直接返回 false
	if len(s)%2 != 0 {
		return false
	}

	// 定义存储左括号的栈
	stack := []rune{}

	// 定义右括号和左括号的映射关系
	rightAndLeftMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 如果map找不到key，返回value的零值，也就是0，通过这个判断是左括号还是右括号
	// 遍历字符串，左括号入栈，遇到第一个右括号，栈顶括号应该是对应的左括号
	for _, v := range s {
		// fmt.Printf("%c\n", v)

		// 能找到key说明是右括号
		if rightAndLeftMap[v] > 0 {
			// 右括号之前全部为0，或者右括号不匹配
			if len(stack) == 0 || rightAndLeftMap[v] != stack[len(stack)-1] {
				return false
			} else {
				// 删除左括号，stack直接赋值为除最后一个元素所有元素
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	// 全部遍历完并且栈为空，就是全部匹配
	return len(stack) == 0
}

func 有效的括号() {
	s := "()[]{}"
	ans := isValid(s)
	println(ans)

}
