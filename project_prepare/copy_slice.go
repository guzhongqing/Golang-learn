package projectprepare

// 这里使用泛型，在使用的时候可以传递任意类型的切片，如果直接使用any，只能传递any类型的切片（any可以接收无类型的字面量，但是不能接收有类型的变量）
func CopySlice[T any](dest, src []T) int {
	// 自己实现和内置copy的功能
	if len(dest) == 0 || len(src) == 0 {
		return 0
	}
	copyNum := len(dest)
	if copyNum > len(src) {
		copyNum = len(src)
	}
	// 复制元素
	for i := 0; i < copyNum; i++ {
		dest[i] = src[i]
	}

	return copyNum

}

// 内置copy
// 如果dest的长度大于src的长度，src会被复制到dest的前copyNum个元素，dest的剩余元素保持不变
// copy函数返回的是复制的元素个数，不是dest的长度
