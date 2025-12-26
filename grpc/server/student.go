package main

import (
	"context"
	"fmt"
	grpc_model "golang_learn/go_basic/grpc/idl/model"
	grpc_service "golang_learn/go_basic/grpc/idl/service"
	"io"
	"math/rand"

	grpc "google.golang.org/grpc"
)

// Student 实现了 StudentServer 接口
type Student struct {
	// 匿名嵌入 → 继承所有方法
	grpc_service.UnimplementedStudentServer
}

func (s Student) QueryStudent(ctx context.Context, query *grpc_service.QueryStudentRequest) (resp *grpc_service.QueryStudentResponse, err error) {
	// 打印请求参数
	fmt.Printf("request: %+v\n", query)
	// 打印上下文信息
	fmt.Printf("context: %+v\n", ctx)
	// 构造响应参数
	resp = &grpc_service.QueryStudentResponse{
		Students: []*grpc_model.Student{
			{Id: 1, Name: "张三", Age: 18},
			{Id: 2, Name: "李四", Age: 23},
		},
	}

	return resp, nil
}

func (s Student) QueryStudents1(ctx context.Context, StudentIds *grpc_service.StudentIds) (resp *grpc_service.QueryStudentResponse, err error) {
	fmt.Printf("QueryStudents1方法")
	return nil, nil
}
func (s Student) QueryStudents2(StudentIds *grpc_service.StudentIds, serverStream grpc.ServerStreamingServer[grpc_model.Student]) error {
	fmt.Printf("QueryStudents2方法")
	for _, id := range StudentIds.Ids {
		stu := &grpc_model.Student{
			Id:   id,
			Name: fmt.Sprintf("姓名%d", id),
			Age:  int32(id)}
		err := serverStream.Send(stu)
		if err != nil {
			fmt.Printf("发送学生%v失败: %+v\n", stu, err)
			return err
		}
	}
	return nil
}
func (s Student) QueryStudents3(clientStream grpc.ClientStreamingServer[grpc_service.StudentId, grpc_service.QueryStudentResponse]) error {
	fmt.Printf("QueryStudents3方法")
	students := make([]*grpc_model.Student, 0, 10)
	for {
		stuId, err := clientStream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端关闭流\n")
				break
			}
			fmt.Printf("接收学生ID失败: %+v\n", err)
			return err
		}
		stu := &grpc_model.Student{
			Id:   stuId.Id,
			Name: fmt.Sprintf("姓名%d", stuId.Id),
			Age:  int32(rand.Intn(100)),
		}
		students = append(students, stu)
	}
	return clientStream.SendMsg(&grpc_service.QueryStudentResponse{
		Students: students,
	})

}
func (s Student) QueryStudents4(bidServer grpc.BidiStreamingServer[grpc_service.StudentId, grpc_model.Student]) error {
	fmt.Printf("QueryStudents4方法")
	for {
		stuId, err := bidServer.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端关闭流\n")
				break
			}
			fmt.Printf("接收学生ID失败: %+v\n", err)
			return err
		}
		stu := &grpc_model.Student{
			Id:   stuId.Id,
			Name: fmt.Sprintf("姓名%d", stuId.Id),
			Age:  int32(rand.Intn(100)),
		}
		err = bidServer.Send(stu)
		if err != nil {
			fmt.Printf("发送学生%v失败: %+v\n", stu, err)
			return err
		}
	}
	return nil
}
