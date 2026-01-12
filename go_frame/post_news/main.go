package main

import (
	"fmt"
	"go_frame/post_news/database/gorm"
	"go_frame/post_news/util"
	"log/slog"
)

func main() {
	fmt.Println("Hello, World!")
	// 初始化配置
	viperConfig := util.InitViper("./conf", "mysql", util.YAML)
	// 初始化数据库连接,相关路径是相对于main.go的路径
	// 2. 初始化数据库连接（带错误捕获）
	if err := gorm.CreateConnection(viperConfig); err != nil {
		slog.Error("数据库初始化失败", "err", err)
	}

}
