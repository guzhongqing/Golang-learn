package main

import (
	"context"
	"fmt"
	grpc_service "golang_learn/go_basic/grpc/idl/service"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

func timer(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Printf("timer调用前\n")
	// 记录开始时间
	start := time.Now()

	// 调用原始方法
	resp, err = handler(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("花费时间: %v ms\n", time.Since(start).Milliseconds())
	fmt.Printf("timer调用后\n")
	return resp, nil

}

func conter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 调用原始方法前，打印方法名
	fmt.Printf("conter调用前\n")
	fmt.Printf("method: %s\n", info.FullMethod)

	// 调用原始方法
	resp, err = handler(ctx, req)
	if err != nil {
		return nil, err
	}
	// 调用原始方法后，打印方法名
	fmt.Printf("conter调用后\n")
	return resp, nil

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	// 监听的地址
	fmt.Printf("server is running on %s\n", listener.Addr().String())
	// 全局设置opts ...ServerOption
	// server := grpc.NewServer(grpc.UnaryInterceptor(timer))
	// 链式调用
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(timer, conter))

	// 注册 StudentServer 服务，可以注册多个service
	grpc_service.RegisterStudentServer(server, Student{})
	server.Serve(listener)
	fmt.Printf("server is stopped\n")

}
