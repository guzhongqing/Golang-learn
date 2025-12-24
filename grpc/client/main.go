package main

import (
	"context"
	"fmt"
	grpc_service "golang_learn/go_basic/grpc/idl/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 创建 StudentClient 客户端
	client := grpc_service.NewStudentClient(conn)
	// 调用 QueryStudent 方法
	req := &grpc_service.QueryStudentRequest{
		Id: 123,
	}
	resp, err := client.QueryStudent(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %+v\n", resp)
}
