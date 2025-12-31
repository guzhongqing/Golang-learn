package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func HttpObservation(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("请求方法：%s\n", r.Method)
	fmt.Printf("请求URL：%s\n", r.URL)
	// 获取请求参数
	fmt.Printf("请求参数：%v\n", r.URL.Query())
	// 获取请求参数,这里Get方法获取的是map[key]切片的第一个值
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Printf("name: %s, age: %s\n", name, age)
	fmt.Printf("请求地址：%s\n", r.Host)
	fmt.Printf("请求协议：%s\n", r.Proto)

	fmt.Printf("请求头：%v\n", r.Header)
	for key, values := range r.Header {
		fmt.Printf("%s: %s\n", key, values)
	}
	fmt.Printf("请求体：\n")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)

	// 返回响应, 包含响应头, 响应状态码, 响应体,设置顺序不能变
	// 设置响应头
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("trace-id", uuid.New().String())
	// 设置响应状态码
	w.WriteHeader(http.StatusNotFound)
	// 写入响应体
	w.Write([]byte(`{"code":404,"msg":"not found"}`))

}

func main() {
	http.HandleFunc("/obs", HttpObservation)
	fmt.Println("server start at 127.0.0.1:8080")

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		panic(err)
	}
}
