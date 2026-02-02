package main

import (
	"go_frame/post_news/database/gorm"
	handler "go_frame/post_news/handler/gin"
	"go_frame/post_news/logger"
	"go_frame/post_news/util"
	"net/http"

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

	// 1. 注册静态资源：路径规范（已符合，无需修改）
	r.Static("/js", "./views/js")
	r.Static("/css", "./views/css")
	// 加载HTML模板：若后续模板有子目录，可改为 ./views/html/**/*
	r.LoadHTMLGlob("./views/html/*")

	// 2. 分组注册：HTML页面路由（前台页面）
	htmlGroup := r.Group("/")
	{
		// 路径改为短横线分隔（kebab-case），状态码使用http标准常量
		htmlGroup.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})
		htmlGroup.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", nil)
		})
		// 路径改为短横线分隔，渲染对应的修改密码模板（需新建modify-password.html）
		htmlGroup.GET("/modify-password", func(c *gin.Context) {
			c.HTML(http.StatusOK, "modify-password.html", nil)
		})
	}

	// 3. 分组注册：API接口路由（统一/api前缀）
	apiGroup := r.Group("/api")
	{
		// 路径补全/前缀，统一为/api/xxx风格，保持命名一致
		apiGroup.POST("/register", handler.RegisterUser)
		apiGroup.POST("/login", handler.LoginUser)
		apiGroup.POST("/logout", handler.LogOutUser)
		// API路径同样使用短横线分隔，与HTML路由风格统一
		apiGroup.POST("/modify-password", handler.ModifyPassword)
	}

	// 启动服务（地址格式规范，无需修改）
	_ = r.Run("localhost:8080")
}
