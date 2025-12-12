package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var filename = "data/test.txt"

func WriteFile() {
	// 写入文件
	if fileOut, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o666); err != nil {
		fmt.Printf("open file %s failed, err: %v\n", filename, err)
	} else {
		defer fileOut.Close()
		fileOut.WriteString("1.这是写入文件的内容\n")
		fileOut.WriteString("2.第二行不换行")
		fileOut.WriteString("3.第三j句内容\n")
		fileOut.WriteString("4.hello world\n")
	}

	// 追加写入文件
	if fileOut, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0o666); err != nil {
		fmt.Printf("open file %s failed, err: %v\n", filename, err)
	} else {
		defer fileOut.Close()
		n, err := fileOut.WriteString("5.这是追加文件的内容\n")
		if err != nil {
			fmt.Printf("write file %s failed, err: %v\n", filename, err)
		} else {
			fmt.Printf("write %d bytes \n", n)
		}

	}
}

func ReadFile() {
	// 读取文件
	if fileIn, err := os.OpenFile(filename, os.O_RDONLY, 0o666); err != nil {
		fmt.Printf("open file %s failed, err: %v\n", filename, err)
	} else {
		defer fileIn.Close()
		buf := make([]byte, 1024)
		fileIn.Seek(0, 0)
		if n, err := fileIn.Read(buf); err != nil {
			fmt.Printf("read file %s failed, err: %v\n", filename, err)
		} else {
			fmt.Printf("read %d bytes: %s\n", n, string(buf))
		}
	}
}

func ReadFileWithBuffer() {
	// 读取文件
	if fileIn, err := os.Open(filename); err != nil {
		fmt.Printf("open file %s failed, err: %v\n", filename, err)
	} else {
		defer fileIn.Close()
		reader := bufio.NewReader(fileIn)
		// 读取文件内容
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				fmt.Printf("read file %s failed, err: %v\n", filename, err)
				// var EOF = errors.New("EOF")
				fmt.Println(io.EOF)
				if err == io.EOF {
					break
				}
			} else {
				fmt.Printf("read line: %s", line)
			}
		}
	}
}

func WriteFileWithBuffer() {
	// 写入文件
	if fileOut, err := os.Create(filename); err != nil {
		fmt.Printf("open file %s failed, err: %v\n", filename, err)
	} else {
		defer fileOut.Close()
		writer := bufio.NewWriter(fileOut)
		n1, _ := writer.WriteString("1.使用bufio写入第一行\n")
		n2, _ := writer.WriteString("2.使用bufio写入第二行\n")
		n3, _ := writer.WriteString("3.使用bufio写入第三行\n")
		n4, _ := writer.WriteString("4.使用bufio写入 hello world\n")
		writer.Flush()
		fmt.Printf("write %d bytes\n", n1+n2+n3+n4)
	}
}
