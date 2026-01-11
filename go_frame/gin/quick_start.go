package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homeHandler(ctx *gin.Context) {
	fmt.Println("homeHandler")
	// 打印请求头
	fmt.Println("请求头：")
	for key, values := range ctx.Request.Header {
		fmt.Println(key, values)
	}
	// 打印请求体
	fmt.Println("请求体：")
	fmt.Println(ctx.Request.Body)

	// 设置响应头
	// ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")// 标准库
	// gin框架设置响应头
	ctx.Header("Content-Type", "application/json")
	// 设置响应体
	// ctx.Writer.WriteHeader(200)           // 标准库设置状态码
	// ctx.Writer.WriteString("hello world") // 标准库设置响应体
	// gin框架设置响应体
	ctx.String(200, "hello world")
	// gin框架设置响应体,这里的gin.H是一个map[string]any的别名,本身就是传入一个map[string]any
	// ctx.JSON(200, map[string]any{
	// 	"message": "hello world",
	// })
	// 使用json序列化
	ctx.JSON(200, gin.H{
		"message": "hello world",
	})

}

// func main() {
// 	engine := gin.Default()
// 	// 注册路由
// 	engine.GET("/home", homeHandler)
// 	// 启动服务
// 	if err := engine.Run("127.0.0.1:8080"); err != nil {
// 		panic(err)
// 	}
// }
