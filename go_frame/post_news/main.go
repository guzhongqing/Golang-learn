package main

import (
	"go_frame/post_news/database/gorm"
	handler "go_frame/post_news/handler/gin"
	"go_frame/post_news/logger"
	"go_frame/post_news/util"

	"github.com/gin-gonic/gin"
)

func Init() {
	// 初始化配置
	// 方式1：传空日志文件名 → 自动生成日期命名的日志文件 2026-01-14_app.log
	// logger.InitLogger("debug", "")

	// 方式2：自定义日志文件名 → 生成 ./log/logger.test.log")
	logger.InitLogger("debug", "logger.test.log")
	defer logger.CloseLogger() // 程序退出释放资源

	viperConfig := util.InitViper("./conf", "mysql", util.YAML)
	// 初始化数据库连接,相关路径是相对于main.go的路径
	// 2. 初始化数据库连接（带错误捕获）
	if err := gorm.CreateConnection(viperConfig); err != nil {
		logger.Error("数据库初始化失败", "err", err)
	}
}

func main() {
	Init()
	r := gin.Default()
	// 注册静态资源，将  ./views/js 目录下的文件映射到 浏览器/js 路径
	r.Static("/js", "./views/js")
	r.Static("/css", "./views/css")
	r.LoadHTMLGlob("./views/html/*")

	// 注册路由
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.LoginUser)
	r.POST("/logout", handler.LogOutUser)
	r.POST("/modify_password", handler.ModifyPassword)

	r.Run("localhost:8080")

}
