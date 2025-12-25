package main

import (
	"context"
	"fmt"
	grpc_service "golang_learn/go_basic/grpc/idl/service"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func timer(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("timer调用前\n")

	// 记录开始时间
	start := time.Now()

	// 调用原始方法
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("花费时间: %v ms\n", time.Since(start).Milliseconds())
	fmt.Printf("timer调用后\n")

	return err
}

func conter(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("conter调用前\n")

	// 调用原始方法
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method: %s\n", method)
	fmt.Printf("conter调用后\n")

	return err
}

func main() {
	// 全局设置opts ...DialOption
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024),
			grpc.MaxCallSendMsgSize(1024),
		),
		// 客户端拦截器
		// grpc.WithUnaryInterceptor(timer),
		// 多个拦截器，链式调用，按顺序执行
		grpc.WithChainUnaryInterceptor(timer, conter),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	const N = 1
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			// 创建 StudentClient 客户端
			client := grpc_service.NewStudentClient(conn)
			// 调用 QueryStudent 方法
			req := grpc_service.QueryStudentRequest{
				Id: 123,
			}
			// 单次调用设置接收消息大小为 1024 字节
			resp, err := client.QueryStudent(context.Background(), &req, grpc.MaxCallRecvMsgSize(1024))
			if err != nil {
				panic(err)
			}
			fmt.Printf("response%d: %+v\n", i, resp)
		}()
	}
	wg.Wait()

}
