package main

import (
	"context"
	"fmt"
	grpc_service "golang_learn/go_basic/grpc/idl/service"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

func counter(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("conter调用前\n")

	// 调用原始方法
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method: %s\n", method)
	fmt.Printf("conter调用后\n")

	return err
}

func devKey(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "dev-key", "123456")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func QueryStudent(client grpc_service.StudentClient, i int) {
	req := grpc_service.QueryStudentRequest{
		Id: 123,
	}
	// 可以单次调用设置接收消息大小为 1024 字节
	resp, err := client.QueryStudent(context.Background(), &req, grpc.MaxCallRecvMsgSize(1024))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response%d: %+v\n", i, resp)

}

func QueryStudent2(client grpc_service.StudentClient, i int) {
	req := grpc_service.StudentIds{
		Ids: []int64{100, 200, 300, 400, 500, 600},
	}

	stream, err := client.QueryStudents2(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("stream EOF结束\n")
				break
			} else {
				fmt.Printf("stream 接收错误: %v\n", err)
				continue
			}
		}
		fmt.Printf("response%d: %+v\n", i, resp)
	}

}

func QueryStudent3(client grpc_service.StudentClient, i int) {
	req := grpc_service.StudentId{
		Id: 999,
	}
	stream, err := client.QueryStudents3(context.Background())
	if err != nil {
		panic(err)
	}
	// 流式发送请求
	for i := 0; i < 6; i++ {
		err := stream.Send(&req)
		if err != nil {
			fmt.Printf("stream 发送错误: %v\n", err)
			continue
		}
		fmt.Printf("stream 发送请求%d成功\n", i)
	}
	// 一次性接收响应
	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("stream 接收错误: %v\n", err)

	}
	for _, student := range resp.Students {
		fmt.Printf("student: %+v\n", student)
	}
}

func QueryStudent4(client grpc_service.StudentClient, i int) {
	stream, err := client.QueryStudents4(context.Background())
	done := make(chan struct{})
	req := grpc_service.StudentId{
		Id: 999,
	}
	if err != nil {
		panic(err)
	}

	// 开启协程接收响应
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Printf("stream EOF结束\n")
					// 因为零缓冲，这里会阻塞，等待其他协程接收完响应，取出元素，解除阻塞
					done <- struct{}{}
					break
				} else {
					fmt.Printf("stream 接收错误: %v\n", err)
					continue
				}
			}
			fmt.Printf("response%d: %+v\n", i, resp)
		}
	}()

	fmt.Println("开始发送")

	//流式发送请求
	for i := 0; i < 6; i++ {
		// 每过一秒发送一次请求
		time.Sleep(time.Second)
		err := stream.Send(&req)
		if err != nil {
			fmt.Printf("stream 发送错误: %v\n", err)
			continue
		}
	}
	// 关闭发送
	stream.CloseSend()
	<-done
	fmt.Printf("stream 接收完成\n")
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
		grpc.WithChainUnaryInterceptor(timer, counter, devKey),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	const N = 1
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			// 创建 StudentClient 客户端
			client := grpc_service.NewStudentClient(conn)

			// 调用 QueryStudents 方法
			// QueryStudent(client, i)
			// 调用 QueryStudent2 方法
			// QueryStudent2(client, i)
			// 调用 QueryStudent3 方法
			// QueryStudent3(client, i)
			// 调用 QueryStudent4 方法
			QueryStudent4(client, i)

		}(i)
	}
	wg.Wait()

}
