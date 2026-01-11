package main

import (
	"bytes"
	"fmt"
	model "go_frame/gin/idl"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func PostPb() {
	fmt.Println("PostPb")

	stu := model.Student{
		Name: "张三",
		Age:  18,
	}
	stuBytes, err := proto.Marshal(&stu)
	if err != nil {
		fmt.Println("marshal failed")
		return
	}
	// 发送http请求
	resp, err := http.Post("http://127.0.0.1:8080/student/BodyBind", "", bytes.NewBuffer(stuBytes))
	if err != nil {
		fmt.Println("post failed")
		return
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed")
		return
	}
	fmt.Println("响应体:", string(body))

}

func main() {
	PostPb()

}
