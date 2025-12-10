package test

import (
	"fmt"
	projectprepare "golang_learn/go_basic/project_prepare"
	"testing"
)

func TestCopySlice(t *testing.T) {

	// 测试用例1：正常情况
	dest1 := []int{1, 2, 3, 4, 5}
	src1 := []int{6, 7, 8, 9, 10}
	copyNum1 := projectprepare.CopySlice(dest1, src1)
	fmt.Printf("dest1: %v\n", dest1)
	if copyNum1 != 5 {
		t.Errorf("copyNum1 != 5, copyNum1 = %d", copyNum1)
	}

	// 测试用例2：源切片为空
	dest2 := []int{1, 2, 3, 4, 5}
	src2 := []int{}
	copyNum2 := projectprepare.CopySlice(dest2, src2)
	fmt.Printf("dest2: %v\n", dest2)
	if copyNum2 != 0 {
		t.Errorf("copyNum2 != 0, copyNum2 = %d", copyNum2)
	}

	// 测试用例3：目标切片为空（目标切片长度为0）
	dest3 := []int{}
	src3 := []int{1, 2, 3, 4, 5}
	copyNum3 := projectprepare.CopySlice(dest3, src3)
	fmt.Printf("dest3: %v\n", dest3)
	if copyNum3 != 0 {
		t.Errorf("copyNum3 != 0, copyNum3 = %d", copyNum3)
	}

	// 测试用例4：目标切片长度小于源切片长度
	dest4 := []int{1, 2, 3, 4}
	src4 := []int{6, 7, 8, 9, 10}
	copyNum4 := projectprepare.CopySlice(dest4, src4)
	fmt.Printf("dest4: %v\n", dest4)
	if copyNum4 != 3 {
		// fmt.Printf("copyNum4 != 3, copyNum4 = %d", copyNum4)
		// t.Fail()
		t.Errorf("copyNum4 != 3, copyNum4 = %d", copyNum4)
		// t.Fatal()
		// t.Fatalf("copyNum4 != 3, copyNum4 = %d", copyNum4)

	}

	// 测试用例5：目标切片长度大于源切片长度
	dest5 := []int{1, 2, 3, 4, 5}
	src5 := []int{6, 7, 8}
	copyNum5 := projectprepare.CopySlice(dest5, src5)
	fmt.Printf("dest5: %v\n", dest5)
	if copyNum5 != 3 {
		t.Errorf("copyNum5 != 3, copyNum5 = %d", copyNum5)
	}

}

// 测试自实现CopySlice函数的性能
func BenchmarkCopySilce(b *testing.B) {
	dest := make([]int, 10_000_000)
	src := make([]int, 10_000_000)
	// 开始计时
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		projectprepare.CopySlice(dest, src)
	}

}

// 测试内置copy()的性能
func BenchmarkStdCopySilce(b *testing.B) {
	dest := make([]int, 10_000_000)
	src := make([]int, 10_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dest, src)
	}

}
