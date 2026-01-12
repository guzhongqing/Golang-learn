package gorm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 初始化GORM日志文件（自动创建目录）
func initGormLogFile() (*os.File, error) {
	// 日志文件路径
	logPath := "../log/gorm.log"
	// 提取日志文件所在的目录路径
	logDir := filepath.Dir(logPath)

	// 检查并创建目录（0755权限：所有者可读可写可执行，其他用户可读可执行）
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err // 目录创建失败时返回错误
	}

	// 打开/创建日志文件（确保目录已存在）
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, // 创建+仅写+追加模式
		os.ModePerm,                         // 系统默认文件权限
	)
	if err != nil {
		return nil, err // 文件创建失败时返回错误
	}

	return logFile, nil
}

func CreateConnection(config *viper.Viper) *gorm.DB {
	host := config.GetString("mysql.host")
	port := config.GetInt("mysql.port")
	username := config.GetString("mysql.username")
	password := config.GetString("mysql.password")
	dbname := config.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// 日志控制
	// 初始化日志文件
	logFile, err := initGormLogFile()
	if err != nil {
		return nil
	}

	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:        500 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:             logger.Info,            // 日志级别
			ParameterizedQueries: false,                  // 为true时，sql日志里面参数使用问号代替
			Colorful:             false,                  // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_", // 表名前缀，`User` 的表名将是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 的表名将是 `user`，而不是默认的 `users`
			// NoLowerCase:   true, // 为true时，字段名不会转换为小写，保持结构体字段的大小写形式，例如 `UserName` 将映射为 `UserName`，而不是默认的 `user_name`
		},
		Logger:               newLogger,
		DryRun:               true,  // 开启DryRun模式，生成SQL日志但不执行
		DisableAutomaticPing: false, // gorm默认会在首次使用DB时执行ping操作，设置为true可以禁用该行为

	})
	if err != nil {
		log.Fatalf("mysql 连接失败: %v", err)
		return nil
	} else {
		log.Printf("mysql 连接成功: %s", dsn)
		// 设置连接池
		// 连接池配置
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("获取数据库连接失败: %v", err)
			return nil
		}
		sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
		sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
		sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期
		return db
	}

}
