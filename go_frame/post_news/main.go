package main

import (
	"fmt"
	"go_frame/post_news/database/gorm"
	"go_frame/post_news/logger"
	"go_frame/post_news/util"
)

func main() {
	fmt.Println("Hello, World!")
	// 初始化配置
	// 方式1：传空日志文件名 → 自动生成日期命名的日志文件 2026-01-14_app.log
	// logger.InitLogger("debug", "")

	// 方式2：自定义日志文件名 → 生成 ./log/user_service.log
	logger.InitLogger("debug", "user_service.log")

	defer logger.CloseLogger() // 程序退出释放资源
	viperConfig := util.InitViper("./conf", "mysql", util.YAML)
	// 初始化数据库连接,相关路径是相对于main.go的路径
	// 2. 初始化数据库连接（带错误捕获）
	if err := gorm.CreateConnection(viperConfig); err != nil {
		logger.Error("数据库初始化失败", "err", err)
	}

}
