package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Scan() {
	fmt.Println("请输入单词：")
	var word string

	fmt.Scanf("%s\n", &word)
	fmt.Printf("你输入的单词是：%s\n", word)

	var num1 int
	var num2 int
	fmt.Println("请输入两个整数：")
	fmt.Scanf("%d %d\n", &num1, &num2)
	fmt.Printf("你输入的两个整数和是：%d\n", num1+num2)

	// 按行读取
	fmt.Println("请输入一行文本：")

	content := make([]byte, 1024)
	n, err := os.Stdin.Read(content)
	if err != nil {
		fmt.Printf("read stdin failed, err: %v\n", err)
	} else {
		fmt.Printf("read %d bytes: %s\n", n, string(content))
	}

}

func Stdout() {
	// 文件句柄不固定
	fmt.Println(os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd())
	fmt.Println("hello world")
	os.Stdout.WriteString("abc")
	os.Stderr.WriteString("213\n")
}

func SysCall() {
	cmd_path, err := exec.LookPath("go")
	if err != nil {
		fmt.Printf("look path failed, err: %v\n", err)
	} else {
		fmt.Printf("go path: %s\n", cmd_path)
	}
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec command failed, err: %v\n", err)
	} else {
		fmt.Printf("go version: %s\n", string(output))
	}
}

func main() {
	// Scan()
	// Stdout()
	 SysCall()
}
