package main

import (
	"fmt"
	grpc_service "golang_learn/go_basic/grpc/idl/service"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	// 监听的地址
	fmt.Printf("server is running on %s\n", listener.Addr().String())
	server := grpc.NewServer()
	// 注册 StudentServer 服务，可以注册多个service
	grpc_service.RegisterStudentServer(server, Student{})
	server.Serve(listener)
	fmt.Printf("server is stopped\n")

}
