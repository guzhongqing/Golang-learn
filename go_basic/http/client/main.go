package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func RequestObservation() {
	// 分隔符
	fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
	// 发送请求
	resp, err := http.Get("http://127.0.0.1:8080/obs?name=张三&age=18")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取响应体
	fmt.Printf("响应体：\n")
	io.Copy(os.Stdout, resp.Body)
	os.Stdout.WriteString("\n")

	// 读取协议
	fmt.Printf("响应协议：%s\n", resp.Proto)
	// 读取状态码
	fmt.Printf("响应状态码：%d\n", resp.StatusCode)
	// 读取状态文本
	fmt.Printf("响应状态文本：%s\n", resp.Status)

	// 读取响应头
	fmt.Printf("响应头：%v\n", resp.Header)
	for key, values := range resp.Header {
		if key == "Date" {
			if tm, err := http.ParseTime(values[0]); err == nil {
				fmt.Printf("%s: %s\n", key, tm.Format("2006-01-02 15:04:05"))
			}

		} else {
			fmt.Printf("%s: %s\n", key, values)

		}
	}

}
func main() {
	RequestObservation()
}
