package test

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

func TestSonic(t *testing.T) {
	fmt.Println("TestSonic")
	sonic.Marshal(2)
}
