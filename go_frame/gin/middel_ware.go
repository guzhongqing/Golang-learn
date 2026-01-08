package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func homeHandler1(ctx *gin.Context) {
	fmt.Println("homeHandler")
	// 打印请求头
	fmt.Println("请求头：")
	for key, values := range ctx.Request.Header {
		fmt.Println(key, values)
	}
	// 打印请求体
	fmt.Println("请求体：")
	defer ctx.Request.Body.Close()
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Printf("ReadAll failed: %v", err)
	}
	fmt.Println(string(body))

	ctx.Header("Content-Type", "application/json")

	ctx.String(200, "hello world")

	ctx.JSON(200, gin.H{
		"message": "hello world",
	})

}

func middleware1(ctx *gin.Context) {
	fmt.Println("middleware1 start")
	// 调用ctx.Abort()后,后面所有的中间件和handler都不会执行
	// ctx.Abort()
	// 调用
	ctx.Next()

	// 后面所有的中间件和handler都执行完成后,才会执行middleware1后面的代码
	// 可以在middleware1后面添加一些代码,这些代码会在所有的中间件和handler执行完成后执行
	fmt.Println("middleware1 end")

}

func main() {
	// gin.Default(),默认使用gin.Logger()和gin.Recovery()中间件
	// gin.Logger()表示每次请求的日志格式为: 日志标识，处理完成时间，HTTP 返回状态码，服务端处理请求耗时，客户端 IP，HTTP 请求方法，请求路径
	// 如，[GIN] 2026/01/08 - 21:38:00 | 200 |       505.7µs |       127.0.0.1 | GET      "/home"
	// gin.Recovery()表示在所有handler中都使用recover()函数,如果handler中发生了panic,则会捕获并打印错误信息
	engine := gin.Default()
	// gin.New(),不使用默认的中间件
	// gin.Default() = gin.New() + gin.Logger() + gin.Recovery()
	// engine := gin.New()
	// engine.Use()添加全局中间件
	// engine.Use(gin.Logger())
	// engine.Use(gin.Recovery())

	// 注册路由
	engine.GET("/home", middleware1, homeHandler1)
	// 启动服务
	if err := engine.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
